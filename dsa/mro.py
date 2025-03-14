class A:
    def method(self):
        print("Method in A")

class B:
    def method(self):
        print("Method in B")

class C():
     def method(self):
        print("Method in C")


class D(C, B, A):
     def method(self):
        print("Method in C")

d = D()
d.method()
