These are **strategies** for handling **conflicts** when multiple threads or processes access shared resources (e.g., data in a database). They are not models but rather **techniques** that can be applied within various concurrency models.

##### **Pessimistic Concurrency Control** 

  **Philosophy**: 
    - Assume conflicts will happen, so prevent them upfront by locking resources.
    - Uses **locks** to prevent other transactions from accessing the same resource.
  **Phases**:
    - Transactions acquire locks before reading or modifying data.
    - Locking overhead
    - Validate -> Read -> Compute -> Write
  
  **Lock Types**:
    - **Shared Locks (Read Locks)** : Allow multiple transactions to read but not modify the data.
    - **Exclusive Locks (Write Locks)** : Allow only one transaction to modify the data.
  
  **Pros**:
    - Prevents conflicts upfront, ensuring consistency and data integrity.
    - Optimized for high-conflict scenarios as it is sequential.
    -  Simple to implement and good for smaller databases.
  
  **Cons**:
    - Can lead to **deadlocks** if locks are not managed properly.
    - Reduces **scalability** and **performance** due to blocking.
  
  **Use Cases**:
- High Contention Systems – Used in banking and ticket booking where conflicts are frequent.
- Critical Data Integrity – Ensures strong consistency in financial transactions and inventory management.
- Relational Databases – Common in SQL databases (e.g., PostgreSQL, MySQL) using row-level locking.
- Long-Running Transactions – Prevents data modification while a transaction is in progress.
- Legacy Systems – Often used in older enterprise applications with strict locking mechanisms.
    

--- 


##### **Optimistic Concurrency Control**

 **Philosophy**:
    - Assumes that conflicts between transactions are rare.

 **Phases**:
    - **Read Phase**: Transactions read data and make modifications locally.
    - **Validation Phase**: Checks if the changes conflict with other transactions.
    - **Write Phase**: If no conflicts are found, the changes are committed; otherwise, the transaction is rolled back and may be retried.
    Read -> Compute -> Validate-> Write

 **Advantages**:
    - Works well in low-conflict scenarios.
    - Reduces locking overhead, improving performance and throughput.
    - Suitable for systems with infrequent updates and for large database or has more records
 **Disadvantages**:
    - High overhead in high-conflict scenarios (contention) due to retries and handling.
    - Requires careful design of the validation process as it is verbose.

 **Example Use Case**:
- Low Contention Systems – Read-heavy applications like analytics dashboards.
- Microservices & Distributed Systems – Ensures consistency without locking in distributed databases.
- Cloud & Multi-Tenant Databases – Used in cloud DBs like DynamoDB, Spanner for concurrency control.
- Offline Sync & Mobile Apps – Ensures correct updates when devices reconnect after being offline.
- Event Sourcing & Blockchain – Prevents conflicts in event-driven and decentralized systems


## Important Links 

https://www.felixcloutier.com/x86/cmpxchg
  



