// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Greeter
 * @dev Store & retrieve value in a variable
 */
contract Greeter {
    constructor() {
    }

    function foo() public returns (bool) {
        A a = new A();
        address(a).call(abi.encodeWithSignature("bar(bool)", true));
        return true;
    }
}

contract A {

    constructor() {
    }

    function bar(bool isRevert) public {
        new B();
        if (isRevert) {
            revert("asdfaewfa");
        }
    }
}

contract B {
    constructor() {
    }
}
