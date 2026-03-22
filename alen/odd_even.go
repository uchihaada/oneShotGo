package main

func oddeven() {

	// // odd and even gorutine using one channel to sync between them but now it will pick randomely one gorutine
	// ch := make(chan (bool), 1)
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func(n int) {
	// 	defer wg.Done()
	// 	for i := 0; i < n; i += 2 {
	// 		<-ch
	// 		fmt.Println(i)
	// 		time.Sleep(500 * time.Millisecond)
	// 		ch <- true
	// 	}
	// }(11)

	// go func(n int) {
	// 	defer wg.Done()
	// 	for i := 1; i < n; i += 2 {
	// 		<-ch
	// 		fmt.Println(i)
	// 		time.Sleep(500 * time.Millisecond)
	// 		ch <- true
	// 	}
	// }(11)

	// ch <- true
	// wg.Wait()

	//  odd and even gorutine using separate channels to sync betwee them it will pick the one that is triggered fast with control
	// evench := make(chan (bool), 1)
	// oddch := make(chan (bool), 1)

	// var wg sync.WaitGroup

	// wg.Add(2)
	// go func(n int) {
	// 	defer wg.Done()
	// 	for i := 0; i < n; i += 2 {
	// 		<-evench
	// 		fmt.Println(i)
	// 		time.Sleep(500 * time.Millisecond)
	// 		oddch <- true
	// 	}
	// }(11)
	// go func(n int) {
	// 	defer wg.Done()
	// 	for i := 1; i < n; i += 2 {
	// 		<-oddch
	// 		fmt.Println(i)
	// 		time.Sleep(500 * time.Millisecond)
	// 		evench <- true
	// 	}
	// }(11)

	// evench <- true
	// wg.Wait()
}
