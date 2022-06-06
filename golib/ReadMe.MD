go build -buildmode=c-shared -o hello.so

```text
golang 生成动态库注意事项:
    - 入参和返回参数使用基础类型
    - 需要调用方实现GoString结构
```

Java GoString:

```
package com.libtest.util;

import com.sun.jna.Structure;

import java.util.ArrayList;
import java.util.List;

public class GoString extends Structure {

    public String str;
    public long length;

    public GoString() {
    }

    public GoString(String str) {
        this.str = str;
        this.length = str.length();
    }

    @Override
    protected List<String> getFieldOrder() {
        List<String> fields = new ArrayList<>();
        fields.add("str");
        fields.add("length");
        return fields;
    }

    public static class ByValue extends GoString implements Structure.ByValue {
        public ByValue() {
        }

        public ByValue(String str) {
            super(str);
        }
    }

    public static class ByReference extends GoString implements Structure.ByReference {
        public ByReference() {
        }

        public ByReference(String str) {
            super(str);
        }
    }
}
```

Java lib:

```
package com.libtest.so;

import com.libtest.util.GoString;
import com.sun.jna.Library;
import com.sun.jna.Native;

public interface LibHello extends Library {
    LibHello INSTANCE = Native.loadLibrary("/home/tan/www/go/src/lib/libhello.so", LibHello.class);

    GoString.ByValue Hello(GoString.ByValue call);
}
```

Java 调用:

```
package com.libtest.controller;

import com.libtest.so.LibHello;
import com.libtest.util.GoString;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
public class TestLibController {
    @Test
    public void test() {
        System.out.println(LibHello.INSTANCE.Hello(new GoString.ByValue("Test")).str);
    }
}
```