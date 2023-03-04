DST := $(patsubst %/interface.go,mock/%_mock.go,$(wildcard */interface.go))

.PHONY: all clean

all: $(DST)

clean:
	rm -f mock/*.go

mock/%_mock.go: %/interface.go
	mockgen -source $^ -destination $@ -package mock
