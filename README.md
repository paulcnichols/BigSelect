An experiment to learn GO
=========================

The goal of this project is to create a thrift server that exposes a queuing interface. Clients may "select" on a set of queues and when messages become available, a "reservation" is returned that represents a lock on the queue.  Any client may "push" into a queue, but clients may "pop" or "peek" into the queue only with a reservation.  
