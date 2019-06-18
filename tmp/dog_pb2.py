# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: dog.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='dog.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\tdog.proto\"%\n\x0fGetMyDogMessage\x12\x12\n\ntarget_dog\x18\x01 \x01(\t\"+\n\rMyDogResponse\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04kind\x18\x02 \x01(\t25\n\x03\x44og\x12.\n\x08GetMyDog\x12\x10.GetMyDogMessage\x1a\x0e.MyDogResponse\"\x00\x62\x06proto3')
)




_GETMYDOGMESSAGE = _descriptor.Descriptor(
  name='GetMyDogMessage',
  full_name='GetMyDogMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='target_dog', full_name='GetMyDogMessage.target_dog', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=13,
  serialized_end=50,
)


_MYDOGRESPONSE = _descriptor.Descriptor(
  name='MyDogResponse',
  full_name='MyDogResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='MyDogResponse.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='kind', full_name='MyDogResponse.kind', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=52,
  serialized_end=95,
)

DESCRIPTOR.message_types_by_name['GetMyDogMessage'] = _GETMYDOGMESSAGE
DESCRIPTOR.message_types_by_name['MyDogResponse'] = _MYDOGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetMyDogMessage = _reflection.GeneratedProtocolMessageType('GetMyDogMessage', (_message.Message,), dict(
  DESCRIPTOR = _GETMYDOGMESSAGE,
  __module__ = 'dog_pb2'
  # @@protoc_insertion_point(class_scope:GetMyDogMessage)
  ))
_sym_db.RegisterMessage(GetMyDogMessage)

MyDogResponse = _reflection.GeneratedProtocolMessageType('MyDogResponse', (_message.Message,), dict(
  DESCRIPTOR = _MYDOGRESPONSE,
  __module__ = 'dog_pb2'
  # @@protoc_insertion_point(class_scope:MyDogResponse)
  ))
_sym_db.RegisterMessage(MyDogResponse)



_DOG = _descriptor.ServiceDescriptor(
  name='Dog',
  full_name='Dog',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=97,
  serialized_end=150,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetMyDog',
    full_name='Dog.GetMyDog',
    index=0,
    containing_service=None,
    input_type=_GETMYDOGMESSAGE,
    output_type=_MYDOGRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DOG)

DESCRIPTOR.services_by_name['Dog'] = _DOG

# @@protoc_insertion_point(module_scope)
