1. 
设置文件名时不要带后缀；
搜索路径可以设置多个，viper 会根据设置顺序依次查找；
viper 获取值时使用section.key的形式，即传入嵌套的键名；
默认值可以调用viper.SetDefault设置。

2.
viper 提供了多种形式的读取方法。在上面的例子中，我们看到了Get方法的用法。Get方法返回一个interface{}的值，使用有所不便。

GetType系列方法可以返回指定类型的值。 其中，Type 可以为Bool/Float64/Int/String/Time/Duration/IntSlice/StringSlice。 但是请注意，如果指定的键不存在或类型不正确，GetType方法返回对应类型的零值。

如果要判断某个键是否存在，使用IsSet方法。 另外，GetStringMap和GetStringMapString直接以 map 返回某个键下面所有的键值对，前者返回map[string]interface{}，后者返回map[string]string。 AllSettings以map[string]interface{}返回所有设置。