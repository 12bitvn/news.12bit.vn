---
title: "JavaScript: in VS hasOwnProperty"
description: "Trong lúc đọc change log của Vuejs thì thấy một commit thay đổi từ dùng in qua hasOwn nên mình cũng tìm hiểu xem nó gây ra lỗi gì và vì sao cần phải thay đổi:
https://github.com/vuejs/vue/commit/733c1be7f5983cdd9e8089a8088b235ba21a4dee Hàm hasOwn mà commit sử dụng chính là TypeScript Wrapper của phương thức hasOwnProperty:
Trước tiên chúng ta thử tìm hiểu document của toán tử in và phương thức hasOwnProperty.
Toán tử &ldquo;in&rdquo; The in operator returns true if the specified property is in the specified object or its prototype chain."
date: 2018-08-15T03:46:30Z
link: "https://12bit.vn/articles/javascript-in-vs-has-own/"
site: site_name
draft: false
---
