// File: src/libraries/IScrollMessenger.sol



pragma solidity ^0.8.16;

interface IScrollMessenger {
    /**********
     * Events *
     **********/

    /// @notice Emitted when a cross domain message is sent.
    /// @param sender The address of the sender who initiates the message.
    /// @param target The address of target contract to call.
    /// @param value The amount of value passed to the target contract.
    /// @param messageNonce The nonce of the message.
    /// @param gasLimit The optional gas limit passed to L1 or L2.
    /// @param message The calldata passed to the target contract.
    event SentMessage(
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 messageNonce,
        uint256 gasLimit,
        bytes message
    );

    /// @notice Emitted when a cross domain message is relayed successfully.
    /// @param messageHash The hash of the message.
    event RelayedMessage(bytes32 indexed messageHash);

    /// @notice Emitted when a cross domain message is failed to relay.
    /// @param messageHash The hash of the message.
    event FailedRelayedMessage(bytes32 indexed messageHash);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice Return the sender of a cross domain message.
    function xDomainMessageSender() external view returns (address);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice Send cross chain message from L1 to L2 or L2 to L1.
    /// @param target The address of account who receive the message.
    /// @param value The amount of ether passed when call target contract.
    /// @param message The content of the message.
    /// @param gasLimit Gas limit required to complete the message relay on corresponding chain.
    function sendMessage(
        address target,
        uint256 value,
        bytes calldata message,
        uint256 gasLimit
    ) external payable;

    /// @notice Send cross chain message from L1 to L2 or L2 to L1.
    /// @param target The address of account who receive the message.
    /// @param value The amount of ether passed when call target contract.
    /// @param message The content of the message.
    /// @param gasLimit Gas limit required to complete the message relay on corresponding chain.
    /// @param refundAddress The address of account who will receive the refunded fee.
    function sendMessage(
        address target,
        uint256 value,
        bytes calldata message,
        uint256 gasLimit,
        address refundAddress
    ) external payable;
}

// File: src/L1/IL1ScrollMessenger.sol



pragma solidity ^0.8.16;

interface IL1ScrollMessenger is IScrollMessenger {
    /**********
     * Events *
     **********/

    /// @notice Emitted when the maximum number of times each message can be replayed is updated.
    /// @param maxReplayTimes The new maximum number of times each message can be replayed.
    event UpdateMaxReplayTimes(uint256 maxReplayTimes);

    /***********
     * Structs *
     ***********/

    struct L2MessageProof {
        // The index of the batch where the message belongs to.
        uint256 batchIndex;
        // Concatenation of merkle proof for withdraw merkle trie.
        bytes merkleProof;
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice Relay a L2 => L1 message with message proof.
    /// @param from The address of the sender of the message.
    /// @param to The address of the recipient of the message.
    /// @param value The msg.value passed to the message call.
    /// @param nonce The nonce of the message to avoid replay attack.
    /// @param message The content of the message.
    /// @param proof The proof used to verify the correctness of the transaction.
    function relayMessageWithProof(
        address from,
        address to,
        uint256 value,
        uint256 nonce,
        bytes memory message,
        L2MessageProof memory proof
    ) external;

    /// @notice Replay an existing message.
    /// @param from The address of the sender of the message.
    /// @param to The address of the recipient of the message.
    /// @param value The msg.value passed to the message call.
    /// @param messageNonce The nonce for the message to replay.
    /// @param message The content of the message.
    /// @param newGasLimit New gas limit to be used for this message.
    /// @param refundAddress The address of account who will receive the refunded fee.
    function replayMessage(
        address from,
        address to,
        uint256 value,
        uint256 messageNonce,
        bytes memory message,
        uint32 newGasLimit,
        address refundAddress
    ) external payable;

    /// @notice Drop a skipped message.
    /// @param from The address of the sender of the message.
    /// @param to The address of the recipient of the message.
    /// @param value The msg.value passed to the message call.
    /// @param messageNonce The nonce for the message to drop.
    /// @param message The content of the message.
    function dropMessage(
        address from,
        address to,
        uint256 value,
        uint256 messageNonce,
        bytes memory message
    ) external;
}
