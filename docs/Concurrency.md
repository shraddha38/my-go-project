
## Introduction to concurrency?

#### Definition?

Concurrency is the ability of a software system to execute multiple tasks or operations at the same time, often on different processors or cores. Software that can do such things are known as _concurrent_ software.

##### Example 
Netflix is **concurrent** as it serves **millions of users simultaneously** by leveraging **distributed systems**, **CDN caching**, and **microservices** - enabling parallel processing and seamless streaming.

####  Why we need a concurrent software / Advantages ? 

1. **Efficiency**: It allows systems to perform multiple tasks simultaneously, saving time and resources. Netflix as a platform can live stream and upload videos. 
2. **Scalability**: It enables systems to handle growing workloads (e.g., more users, requests, or data) without slowing down. Multiple videos being uploaded in different countries
3. **User Experience**: It ensures smooth and uninterrupted experiences for users, especially in high-demand applications like Netflix or online banking.
4. **Fault Tolerance**: It allows systems to continue functioning even if some tasks fail, ensuring reliability. Netflix video streaming will continue to work even if the download feature breaks.

---

### Challenges of Concurrency 

#### 1. Race Conditions 
Simply put, a race condition is a flaw in concurrent programming that occurs when multiple processes or threads of a program attempt to access a resource at the same time. In the context of databases, it is a condition where multiple database transactions or operations try to access and modify the same data concurrently.

#####  Example
1. **User A → Bank App:** Request withdraw $90
2. **User B → Bank App:** Request withdraw $90 (simultaneous)
3. **Bank App → Database:** Read balance ($100) (twice)
4. **Database → Bank App:** Returns balance $100 (to both requests)
5. **Bank App → Database:** Deducts $90 (twice, leading to incorrect -$80 balance)
6. **Database → Bank App:** Approves both transactions
7. **Bank App → User A & User B:** Withdrawal successful

This would be a loss for the bank, and a very dangerous loophole in the application.

#### 2. Deadlocks
A **deadlock** occurs when two or more processes are stuck in a cycle, each waiting for a resource that the other is holding, preventing further execution.
#####  Example
Consider two bank transactions happening at the same time:
1. **Transaction A** locks **Account X** and waits for **Account Y**.
2. **Transaction B** locks **Account Y** and waits for **Account X**.
3. Neither transaction can proceed because they are waiting for each other, causing a **deadlock**.

#### 3. Data Inconsistency and Performance Overhead

Without proper synchronization (coordinating the execution of threads and processes) , shared data can become **corrupt** or inconsistent.
Synchronization mechanisms (locks) introduce **waiting time**, reducing efficiency.

---

### What is pessimistic vs optimistic concurrency control? 
 - Are they locks or strategies ?
##### Answer- 
Optimistic and pessimistic are two distinct strategies for handling **concurrency control** in databases and distributed systems. They are **not the same as locks**, but locks are often used as a part of the implementation. 



