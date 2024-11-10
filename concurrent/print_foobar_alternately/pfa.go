package print_foobar_alternately

type FooBar struct {
	n        int
	fch, bch chan struct{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:   n,
		fch: make(chan struct{}, 1),
		bch: make(chan struct{}, 1),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.fch <- struct{}{}
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.bch <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.bch
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		<-fb.fch
	}
}
