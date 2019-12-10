goroutine이란 
===============

Go루틴은 Go Runtime 이 관리하는 Logical Thread이다. Go에서 "go" 키워드를 사용해 함수를 호출하면, Runtime시 새로운 goroutine을 실행한다.

goroutine은 asynchronusly하게 routine을 실행한다. 따라서 여러 코드를 **concurrent**하게 실행하는데 사용된다.

goroutine은 OS Thread 보다 훨씬 가볍게 Async Cocurrent 처리를 구현하기 위해 만든 것이다. 이는 기본적으로 Go Runtime이 자체 관리한다.

Go Runtime 상에서 관리된느 작업단위인 여러 goroutine은 종종 하나의 OS Thread 1개로도 실행되곤 한다.

즉, goroutine은 OS Thread와 1대1로 매핑되지 않고, Multiplexing으로 훨씬 적은 OS Thread를 사용한다. 

Memory측면에서도 일반 OS Thread가 1MB의 stack을 갖는 반면, goroutine은 이보다 훨씬 작은 몇 KB의 stack을 갖는다고 한다. 또한 이는 필요시 동적으로 증가한다. 

**Go Runtime은 goroutine을 관리하면서 Go Channel을 통해 goroutine간 통신을 쉽게 할 수 있도록 한다.**
