package main

import "testing"

func TestGreetingSpecificJohn(t *testing.T) {
	greeting := CreateGreeting("John")
	if greeting != "Hello, John\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificDemo(t *testing.T) {
	greeting := CreateGreeting("Demo")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestGreetingSpecificPaul(t *testing.T) {
	greeting := CreateGreeting("Paul")
	if greeting != "Hello, Paul\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestGreetingSpecificPaul2(t *testing.T) {
	greeting := CreateGreeting("Paul2")
	if greeting != "Hello, Paul2\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestGreetingSpecificPaul3(t *testing.T) {
	greeting := CreateGreeting("Paul3")
	if greeting != "Hello, Paul3\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestGreetingSpecificPaul4(t *testing.T) {
	greeting := CreateGreeting("Paul4")
	if greeting != "Hello, Paul4\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestGreetingSpecificPaul5(t *testing.T) {
	greeting := CreateGreeting("Paul5")
	if greeting != "Hello, Paul5\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestGreetingSpecificPaul6(t *testing.T) {
	greeting := CreateGreeting("Paul6")
	if greeting != "Hello, Paul6\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestGreetingSpecificPaul7(t *testing.T) {
	greeting := CreateGreeting("Paul7")
	if greeting != "Hello, Paul7\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestShowFailure(t *testing.T) {
	greeting := CreateGreeting("Demo1")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Intentional failure. got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}



func TestGreetingDefault(t *testing.T) {
	greeting := CreateGreeting("")
	if greeting != "Hello, Guest\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
	}
}

func TestShowFailure2(t *testing.T) {
	greeting := CreateGreeting("Demo1")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Intentional failure. got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}
