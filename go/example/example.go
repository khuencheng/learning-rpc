// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package example

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - Text
type Data struct {
  Text string `thrift:"text,1" db:"text" json:"text"`
}

func NewData() *Data {
  return &Data{}
}


func (p *Data) GetText() string {
  return p.Text
}
func (p *Data) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *Data)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Text = v
}
  return nil
}

func (p *Data) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Data"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *Data) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("text", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:text: ", p), err) }
  if err := oprot.WriteString(string(p.Text)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.text (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:text: ", p), err) }
  return err
}

func (p *Data) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Data(%+v)", *p)
}

type FormatData interface {
  // Parameters:
  //  - Data
  ArtificialIdiot(ctx context.Context, data *Data) (r *Data, err error)
}

type FormatDataClient struct {
  c thrift.TClient
}

func NewFormatDataClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *FormatDataClient {
  return &FormatDataClient{
    c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
  }
}

func NewFormatDataClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *FormatDataClient {
  return &FormatDataClient{
    c: thrift.NewTStandardClient(iprot, oprot),
  }
}

func NewFormatDataClient(c thrift.TClient) *FormatDataClient {
  return &FormatDataClient{
    c: c,
  }
}

func (p *FormatDataClient) Client_() thrift.TClient {
  return p.c
}
// Parameters:
//  - Data
func (p *FormatDataClient) ArtificialIdiot(ctx context.Context, data *Data) (r *Data, err error) {
  var _args0 FormatDataArtificialIdiotArgs
  _args0.Data = data
  var _result1 FormatDataArtificialIdiotResult
  if err = p.Client_().Call(ctx, "artificial_idiot", &_args0, &_result1); err != nil {
    return
  }
  return _result1.GetSuccess(), nil
}

type FormatDataProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler FormatData
}

func (p *FormatDataProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *FormatDataProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *FormatDataProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewFormatDataProcessor(handler FormatData) *FormatDataProcessor {

  self2 := &FormatDataProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self2.processorMap["artificial_idiot"] = &formatDataProcessorArtificialIdiot{handler:handler}
return self2
}

func (p *FormatDataProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(ctx, seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x3.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush(ctx)
  return false, x3

}

type formatDataProcessorArtificialIdiot struct {
  handler FormatData
}

func (p *formatDataProcessorArtificialIdiot) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := FormatDataArtificialIdiotArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("artificial_idiot", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return false, err
  }

  iprot.ReadMessageEnd()
  result := FormatDataArtificialIdiotResult{}
var retval *Data
  var err2 error
  if retval, err2 = p.handler.ArtificialIdiot(ctx, args.Data); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing artificial_idiot: " + err2.Error())
    oprot.WriteMessageBegin("artificial_idiot", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush(ctx)
    return true, err2
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("artificial_idiot", thrift.REPLY, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
    err = err2
  }
  if err != nil {
    return
  }
  return true, err
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Data
type FormatDataArtificialIdiotArgs struct {
  Data *Data `thrift:"data,1" db:"data" json:"data"`
}

func NewFormatDataArtificialIdiotArgs() *FormatDataArtificialIdiotArgs {
  return &FormatDataArtificialIdiotArgs{}
}

var FormatDataArtificialIdiotArgs_Data_DEFAULT *Data
func (p *FormatDataArtificialIdiotArgs) GetData() *Data {
  if !p.IsSetData() {
    return FormatDataArtificialIdiotArgs_Data_DEFAULT
  }
return p.Data
}
func (p *FormatDataArtificialIdiotArgs) IsSetData() bool {
  return p.Data != nil
}

func (p *FormatDataArtificialIdiotArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FormatDataArtificialIdiotArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.Data = &Data{}
  if err := p.Data.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Data), err)
  }
  return nil
}

func (p *FormatDataArtificialIdiotArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("artificial_idiot_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FormatDataArtificialIdiotArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("data", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:data: ", p), err) }
  if err := p.Data.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Data), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:data: ", p), err) }
  return err
}

func (p *FormatDataArtificialIdiotArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FormatDataArtificialIdiotArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FormatDataArtificialIdiotResult struct {
  Success *Data `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewFormatDataArtificialIdiotResult() *FormatDataArtificialIdiotResult {
  return &FormatDataArtificialIdiotResult{}
}

var FormatDataArtificialIdiotResult_Success_DEFAULT *Data
func (p *FormatDataArtificialIdiotResult) GetSuccess() *Data {
  if !p.IsSetSuccess() {
    return FormatDataArtificialIdiotResult_Success_DEFAULT
  }
return p.Success
}
func (p *FormatDataArtificialIdiotResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *FormatDataArtificialIdiotResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField0(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FormatDataArtificialIdiotResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Data{}
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *FormatDataArtificialIdiotResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("artificial_idiot_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField0(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FormatDataArtificialIdiotResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *FormatDataArtificialIdiotResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FormatDataArtificialIdiotResult(%+v)", *p)
}


