
// combination of queue and server generated secret represent a locked queue
struct Reservation {
	1: required string queue,
	2: required string secret
}

service BigSelect {

	// returns a log on a given queue
    Reservation select(1: list<string> queues),
    void release(1: Reservation reservation),
    
    void push(1: string queue, 2: string data),
    string pop(1: Reservation reservation),
    list<string> peek(1: Reservation reservation, 2: i32 nitems = -1)
}
