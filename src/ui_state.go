package main

var ui_state UIState

type UIState struct {
    Boards              BoardsState
    Posts               PostsState
    Threads             ThreadsState
}

func (s *UIState) LockViews() {
    s.Boards.WaitGroup.Add(1)
    s.Posts.WaitGroup.Add(1)
    s.Threads.WaitGroup.Add(1)
}

func (s *UIState) WaitUntilInitialized() {
    s.Boards.WaitGroup.Wait()
    s.Posts.WaitGroup.Wait()
    s.Threads.WaitGroup.Wait()
}
