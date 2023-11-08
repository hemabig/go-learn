1. 参数
    Get /api/:user/age

2. 防止冲突， 冲突会panic
   eg: 冲突
    Get /api/:user/age
    Get /api/info

    不冲突
    Get /api/:user/age
    POST /api/info

3. 参数最多 255

4. * 不能再路径的开头