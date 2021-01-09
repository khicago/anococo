package anococo

// exampleOfTypeDefAnococo
// <% memo: this is a sample for typedef parser %>
// <% btw: if you write a doc site generator, this may be useful %>
// <% fyi: you can run command below to test it by yourself %>
// <% e.g. go run ./cmd/anococo t %>
type exampleOfTypeDefAnococo1 struct{}

// exampleOfTypeDefAnococo batch
// <% memo: this is another sample %>
type (
	exampleOfTypeDefAnococo2 struct{}

	// <% e.g.: this is sample 3 %>
	exampleOfTypeDefAnococo3 struct{}
)
