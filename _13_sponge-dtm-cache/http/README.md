## Cache Consistency

Using a web service created by [Sponge](https://github.com/zhufuyi/sponge) combined with [DTM](https://github.com/dtm-labs/dtm) and [RocksCache](https://github.com/dtm-labs/rockscache), this example demonstrates cache consistency (with Redis and MySQL) including `eventual consistency`, `atomicity`, `strong consistency`, and `strong consistency during downgrade and upgrade`.

<br>

### Quick Start

- Start the Redis service.
- Start the MySQL service and import the [stock.sql](test/stock.sql) file into the database.
- Download the [DTM](https://github.com/dtm-labs/dtm/releases/tag/v1.18.0) executable, modify the default DTM configuration to use Redis, then start the DTM service with: `dtm -c conf.yml`.
- Clone the project code locally and modify the IP addresses in the MySQL, Redis, and DTM configurations in [config.yml](configs/stock.yml) (replace the default IP addresses 192.168.3.37 and 192.168.3.90).

Compile and start the service:

```bash
make run
```

Open [http://localhost:8080/apis/swagger/index.html](http://localhost:8080/apis/swagger/index.html) in your browser to test the four different cache consistency approaches.

![cache-http-pb-swagger](https://raw.githubusercontent.com/zhufuyi/sponge_examples/main/assets/cache-http-pb-swagger.png)

#### Eventual Consistency

Using the "mark deletion" strategy for the cache thoroughly solves the inconsistency between the database and the cache that is not resolved by merely deleting the cache, ensuring eventual consistency even under extreme conditions.

**Example code: [final.go](internal/handler/final.go).**

<br>

#### Atomicity

This approach ensures that even in the event of a process crash, updates to the database and the cache are either both successful or both fail, making it simpler than other architectures like local message tables, transaction messages, or binlog listeners.

**Example code: [atomic.go](internal/handler/atomic.go).**

<br>

#### Strong Consistency

The prerequisite for strong consistency is that "all data reads must be from the cache". For both the database and Redis, if all reads are provided only by the cache, strong consistency can be achieved easily. The `Fetch` function in RocksCache offers strong consistent cache reads by not returning outdated data and instead synchronously waits for the latest result.

For example, in a recharge scenario, after a user successfully recharges, if the user queries the business result (by checking whether the two-phase global transaction status has succeeded), the system will inform the user of the incomplete status until the global transaction completes, even if the database has been updated.

Strong consistency comes with a cost, mainly performance degradation. Compared to eventual consistency, strong consistency in data reads requires waiting for the latest results, which increases response latency, and also may involve waiting for results from other processes, which consumes resources.

**Example code: [strong.go](internal/handler/strong.go).**

<br>

#### Strong Consistency During Downgrade and Upgrade

Downgrade refers to reading data from the database when the cache is faulty, while upgrade refers to reading data from the cache after the cache recovers. During the short time window of downgrading and upgrading, strong consistency can still be maintained.

- Use DTM's Saga mode to update data, ensuring atomicity of the three operations: locking the cache, updating the database, and deleting the cache.
- After updating the database but before updating the cache, the system can inform the user that the business is complete (unlike the earlier strong consistency scenario, this condition is relaxed).
- In strong consistency access mode, queries will wait for the data update result.
- During downgrade, first disable cache reads and wait until no read operations access the cache, then disable cache deletion.
- During upgrade, first enable cache deletion to ensure all database updates are reflected in the cache, then enable cache reads.

**Example code: [downgrade.go](internal/handler/downgrade.go).**

<br>

Reference:

- https://dtm.pub/app/cache.html
- https://github.com/dtm-labs/dtm-cases/tree/main/cache
