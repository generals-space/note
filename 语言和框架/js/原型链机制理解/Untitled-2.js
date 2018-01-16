var _get = function get(object, property, receiver) { 
    if (object === null) object = Function.prototype; 
    var desc = Object.getOwnPropertyDescriptor(object, property); 
    if (desc === undefined) { 
        var parent = Object.getPrototypeOf(object); 
        if (parent === null) { return undefined; } 
        else { return get(parent, property, receiver); } 
    } else if ("value" in desc) { 
        return desc.value; 
    } else { 
        var getter = desc.get; 
        if (getter === undefined) {
             return undefined; 
        } 
        return getter.call(receiver); 
    } 
};

var _createClass = function () { 
    function defineProperties(target, props) { 
        for (var i = 0; i < props.length; i++) { 
            var descriptor = props[i]; 
            descriptor.enumerable = descriptor.enumerable || false; 
            descriptor.configurable = true; 
            if ("value" in descriptor) 
                descriptor.writable = true; 
            Object.defineProperty(target, descriptor.key, descriptor); 
        } 
    } 
    return function (Constructor, protoProps, staticProps) { 
        if (protoProps) defineProperties(Constructor.prototype, protoProps); 
        if (staticProps) defineProperties(Constructor, staticProps); 
        return Constructor; 
    }; 
}();

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

// 貌似不支持多重继承啊.
function _inherits(subClass, superClass) {
    if (typeof superClass !== "function" && superClass !== null) { 
        throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); 
    } 
    // 覆写子类的prototype对象
    subClass.prototype = Object.create(superClass && superClass.prototype, { 
        constructor: { 
            value: subClass, 
            enumerable: false, 
            writable: true, 
            configurable: true 
        } 
    }); 
    // 设置隐式原型, 感觉这样很怪. 因为这样意为着子类将成为父类的实例对象...呃, 类似的概念
    // 但我不觉得父子类关系与类和实例的关系一样...
    if (superClass) 
        Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; 
}

function _classCallCheck(instance, Constructor) {
    if (!(instance instanceof Constructor)) { 
        throw new TypeError("Cannot call a class as a function"); 
    } 
}

var A = function () {
    function A() {
        // 检查当前this对象是否为A的实例, 如果不是说明是当成函数直接用的...
        _classCallCheck(this, A);
    }
    // 创建类属性
    _createClass(A, [{
        key: "render",
        value: function render() {
            console.log(1);
        }
    }]);

    return A;
}();

var B = function (_A) {
    _inherits(B, _A);

    function B() {
        _classCallCheck(this, B);

        return _possibleConstructorReturn(this, (B.__proto__ || Object.getPrototypeOf(B)).call(this));
    }
    // 原型方法中取到了构造函数类本身, 感觉这样耦合性比较大
    // 这是直接到`B.prototype.__proto__`指向的原型链上寻找目标方法, 
    // 但我不想每次在写子类方法时还要显示写父类变量.
    _createClass(B, [{
        key: "render",
        value: function render() {
            _get(B.prototype.__proto__ || Object.getPrototypeOf(B.prototype), "render", this).call(this);
            console.log(2);
        }
    }]);

    return B;
}(A);