仔细想了想, maven/gradle 的作用就类似于 python 中的 pip + setuptools, ta不只能实现依赖管理, 也能创建基本的代码框架, 实现整个工程的构建, 打包与运行.

pom.xml(Project Object Model) 是 maven 依赖的配置文件, 由于 maven 比单纯的 pip 功能强大, 配置也更复杂, 从某种程度来说, pom.xml 类似于 `requirements.txt` + `setuptools`的`setup.py` + 模块管理中的`__init__.py` + ...的功能.

相比于 C/C++ 中的 makefile, maven(pip + setuptools) 指定了工程目录的结构, 是写死的, 依赖包的存储路径也是规定好的; 好处是不需要像 makefile 手工写一大串 build/test 的目标命令.

其实要比较的话, 应该和 node 下的编译工具链更相似一点, npm, webpack 等, 不过我对这些也不是太了解...
