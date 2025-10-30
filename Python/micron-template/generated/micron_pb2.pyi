from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from collections.abc import Mapping as _Mapping
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class MessageRequest(_message.Message):
    __slots__ = ("command", "payload")
    COMMAND_FIELD_NUMBER: _ClassVar[int]
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    command: str
    payload: str
    def __init__(self, command: _Optional[str] = ..., payload: _Optional[str] = ...) -> None: ...

class MessageReply(_message.Message):
    __slots__ = ("payload", "status")
    class Status(_message.Message):
        __slots__ = ("code", "is_error", "desc")
        CODE_FIELD_NUMBER: _ClassVar[int]
        IS_ERROR_FIELD_NUMBER: _ClassVar[int]
        DESC_FIELD_NUMBER: _ClassVar[int]
        code: int
        is_error: bool
        desc: str
        def __init__(self, code: _Optional[int] = ..., is_error: bool = ..., desc: _Optional[str] = ...) -> None: ...
    PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    payload: str
    status: MessageReply.Status
    def __init__(self, payload: _Optional[str] = ..., status: _Optional[_Union[MessageReply.Status, _Mapping]] = ...) -> None: ...
