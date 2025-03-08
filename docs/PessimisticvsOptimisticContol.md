These are **strategies** for handling **conflicts** when multiple threads or processes access shared resources (e.g., data in a database). They are not models but rather **techniques** that can be applied within various concurrency models.

##### **Pessimistic Concurrency Control** 

- **Philosophy**: 
  . - Assume conflicts will happen, so prevent them upfront by locking resources.
    - Uses **locks** to prevent other transactions from accessing the same resource.
    - Transactions acquire locks before reading or modifying data.
    - Locking overhead
    - Validate -> Read -> Compute -> Write
2. **Lock Types**:
    - **Shared Locks (Read Locks)** : Allow multiple transactions to read but not modify the data.
    - **Exclusive Locks (Write Locks)** : Allow only one transaction to modify the data.
3. **Advantages**:
    - Prevents conflicts upfront, ensuring consistency.
    - Optimized for high-conflict scenarios.
- **Use Cases**:
    - Systems where conflicts are frequent and critical (e.g., banking systems, reservation systems,  ticketmaster ).
    - Scenarios where data integrity is more important than performance.
- **Pros**:
    - Ensures no conflicts occur during the operation.
    - Simple to implement and good for smaller databases.
- **Cons**:
    - Can lead to **deadlocks** if locks are not managed properly.
    - Reduces **scalability** and **performance** due to blocking.
- **Example**: A ticket booking system locks a seat while a user is in the process of purchasing it.

##### **Optimistic Concurrency Control (OCC)**

1. **Philosophy**:
    - Assumes that conflicts between transactions are rare.
    - Transactions are allowed to proceed without locking resources initially.
    - Validation is performed at the **commit time** to check for conflicts.
    - Rollback and retry overhead
2. **Phases**:
    - **Read Phase**: Transactions read data and make modifications locally.
    - **Validation Phase**: Checks if the changes conflict with other transactions.
    - **Write Phase**: If no conflicts are found, the changes are committed; otherwise, the transaction is rolled back and may be retried.
    Read -> Compute -> Validate-> Write
3. **Advantages**:
    - Works well in low-conflict scenarios.
    - Reduces locking overhead, improving performance.
    - Suitable for systems with infrequent updates and for large database or has more records
4. **Disadvantages**:
    - High overhead in high-conflict scenarios due to rollbacks and retries.
    - Requires careful design of the validation process.
5. **Example Use Case**:

