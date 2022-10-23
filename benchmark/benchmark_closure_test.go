package benchmark

import (
	"github.com/crossoverJie/gscript"
	"testing"
)

func BenchmarkClosure(b *testing.B) {
	script := `class ListNode{
    int value;
    ListNode next;
    ListNode(int v, ListNode n){
        value =v;
        next = n;
    }
}

// 两个对象比较需要实现运算符重载
bool operator == (ListNode p1, ListNode p2){
    return p1.value == p2.value;
}

bool hasCycle(ListNode head){
    if (head == nil){
        return false;
    }
    if (head.next == nil){
        return false;
    }

    ListNode fast = head.next;
    ListNode slow = head;
    // bool ret = false;
    for (fast.next != nil){

        if (fast == slow){
            return true;
        }

        if (fast.next == nil){
            return false;
        }
        if (fast.next.next == nil){
            return false;
        }
        if (slow.next == nil){
            return false;
        }

        fast = fast.next.next;
        slow = slow.next;
    }
    return false;
}

ListNode l1 = ListNode(1, nil);
bool b1 =hasCycle(l1);
println(b1);
assertEqual(b1, false);

ListNode l4 = ListNode(4, nil);
ListNode l3 = ListNode(3, l4);
ListNode l2 = ListNode(2, l3);
bool b2 = hasCycle(l2);
println(b2);
assertEqual(b2, false);

l4.next = l2;
bool b3 = hasCycle(l2);
println(b3);
assertEqual(b3, true);`
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		gscript.NewCompiler().Compiler(script)
	}
}
