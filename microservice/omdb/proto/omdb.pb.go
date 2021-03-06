// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/omdb.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Page    int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_omdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_proto_omdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_proto_omdb_proto_rawDescGZIP(), []int{0}
}

func (x *Param) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *Param) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type MovieResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbId string `protobuf:"bytes,3,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *MovieResult) Reset() {
	*x = MovieResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_omdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieResult) ProtoMessage() {}

func (x *MovieResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_omdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieResult.ProtoReflect.Descriptor instead.
func (*MovieResult) Descriptor() ([]byte, []int) {
	return file_proto_omdb_proto_rawDescGZIP(), []int{1}
}

func (x *MovieResult) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieResult) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieResult) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *MovieResult) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieResult) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type SearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search       []*MovieResult `protobuf:"bytes,1,rep,name=search,proto3" json:"search,omitempty"`
	TotalResults string         `protobuf:"bytes,2,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	Response     string         `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	Error        string         `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SearchResult) Reset() {
	*x = SearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_omdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResult) ProtoMessage() {}

func (x *SearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_omdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResult.ProtoReflect.Descriptor instead.
func (*SearchResult) Descriptor() ([]byte, []int) {
	return file_proto_omdb_proto_rawDescGZIP(), []int{2}
}

func (x *SearchResult) GetSearch() []*MovieResult {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *SearchResult) GetTotalResults() string {
	if x != nil {
		return x.TotalResults
	}
	return ""
}

func (x *SearchResult) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *SearchResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_omdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_proto_omdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_proto_omdb_proto_rawDescGZIP(), []int{3}
}

func (x *Rating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Rating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MovieDetailResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year       string    `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Rated      string    `protobuf:"bytes,3,opt,name=rated,proto3" json:"rated,omitempty"`
	Released   string    `protobuf:"bytes,4,opt,name=released,proto3" json:"released,omitempty"`
	Runtime    string    `protobuf:"bytes,5,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Genre      string    `protobuf:"bytes,6,opt,name=genre,proto3" json:"genre,omitempty"`
	Director   string    `protobuf:"bytes,7,opt,name=director,proto3" json:"director,omitempty"`
	Writer     string    `protobuf:"bytes,8,opt,name=writer,proto3" json:"writer,omitempty"`
	Actors     string    `protobuf:"bytes,9,opt,name=actors,proto3" json:"actors,omitempty"`
	Plot       string    `protobuf:"bytes,10,opt,name=plot,proto3" json:"plot,omitempty"`
	Language   string    `protobuf:"bytes,11,opt,name=language,proto3" json:"language,omitempty"`
	Country    string    `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	Awards     string    `protobuf:"bytes,13,opt,name=awards,proto3" json:"awards,omitempty"`
	Poster     string    `protobuf:"bytes,14,opt,name=poster,proto3" json:"poster,omitempty"`
	Ratings    []*Rating `protobuf:"bytes,15,rep,name=ratings,proto3" json:"ratings,omitempty"`
	Metascore  string    `protobuf:"bytes,16,opt,name=metascore,proto3" json:"metascore,omitempty"`
	ImdbRating string    `protobuf:"bytes,17,opt,name=imdb_rating,json=imdbRating,proto3" json:"imdb_rating,omitempty"`
	ImdbVotes  string    `protobuf:"bytes,18,opt,name=imdb_votes,json=imdbVotes,proto3" json:"imdb_votes,omitempty"`
	ImdbId     string    `protobuf:"bytes,19,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	Type       string    `protobuf:"bytes,20,opt,name=type,proto3" json:"type,omitempty"`
	Dvd        string    `protobuf:"bytes,21,opt,name=dvd,proto3" json:"dvd,omitempty"`
	BoxOffice  string    `protobuf:"bytes,22,opt,name=box_office,json=boxOffice,proto3" json:"box_office,omitempty"`
	Production string    `protobuf:"bytes,23,opt,name=production,proto3" json:"production,omitempty"`
	Website    string    `protobuf:"bytes,24,opt,name=website,proto3" json:"website,omitempty"`
	Response   string    `protobuf:"bytes,25,opt,name=response,proto3" json:"response,omitempty"`
	Error      string    `protobuf:"bytes,26,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MovieDetailResult) Reset() {
	*x = MovieDetailResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_omdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetailResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetailResult) ProtoMessage() {}

func (x *MovieDetailResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_omdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetailResult.ProtoReflect.Descriptor instead.
func (*MovieDetailResult) Descriptor() ([]byte, []int) {
	return file_proto_omdb_proto_rawDescGZIP(), []int{4}
}

func (x *MovieDetailResult) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieDetailResult) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieDetailResult) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *MovieDetailResult) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *MovieDetailResult) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *MovieDetailResult) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *MovieDetailResult) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieDetailResult) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *MovieDetailResult) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *MovieDetailResult) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *MovieDetailResult) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *MovieDetailResult) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MovieDetailResult) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *MovieDetailResult) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *MovieDetailResult) GetRatings() []*Rating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *MovieDetailResult) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *MovieDetailResult) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *MovieDetailResult) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *MovieDetailResult) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *MovieDetailResult) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieDetailResult) GetDvd() string {
	if x != nil {
		return x.Dvd
	}
	return ""
}

func (x *MovieDetailResult) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *MovieDetailResult) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *MovieDetailResult) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *MovieDetailResult) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *MovieDetailResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_omdb_proto protoreflect.FileDescriptor

var file_proto_omdb_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6f, 0x6d, 0x64, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x7c, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x64, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x90, 0x01,
	0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x29,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x36, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb5, 0x05, 0x0a, 0x11, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x74,
	0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x64,
	0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x64, 0x62, 0x5f,
	0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x64,
	0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x76, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x64, 0x76, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x78, 0x5f, 0x6f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x78, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x7a, 0x0a, 0x0b, 0x4f, 0x4d, 0x44, 0x42, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0b, 0x2e, 0x6f, 0x6d, 0x64, 0x62,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x12, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x17, 0x2e, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x6f, 0x6d, 0x64, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_omdb_proto_rawDescOnce sync.Once
	file_proto_omdb_proto_rawDescData = file_proto_omdb_proto_rawDesc
)

func file_proto_omdb_proto_rawDescGZIP() []byte {
	file_proto_omdb_proto_rawDescOnce.Do(func() {
		file_proto_omdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_omdb_proto_rawDescData)
	})
	return file_proto_omdb_proto_rawDescData
}

var file_proto_omdb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_omdb_proto_goTypes = []interface{}{
	(*Param)(nil),                  // 0: omdb.Param
	(*MovieResult)(nil),            // 1: omdb.MovieResult
	(*SearchResult)(nil),           // 2: omdb.SearchResult
	(*Rating)(nil),                 // 3: omdb.Rating
	(*MovieDetailResult)(nil),      // 4: omdb.MovieDetailResult
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_proto_omdb_proto_depIdxs = []int32{
	1, // 0: omdb.SearchResult.search:type_name -> omdb.MovieResult
	3, // 1: omdb.MovieDetailResult.ratings:type_name -> omdb.Rating
	0, // 2: omdb.OMDBService.Search:input_type -> omdb.Param
	5, // 3: omdb.OMDBService.Get:input_type -> google.protobuf.StringValue
	2, // 4: omdb.OMDBService.Search:output_type -> omdb.SearchResult
	4, // 5: omdb.OMDBService.Get:output_type -> omdb.MovieDetailResult
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_omdb_proto_init() }
func file_proto_omdb_proto_init() {
	if File_proto_omdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_omdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_omdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_omdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_omdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_omdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetailResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_omdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_omdb_proto_goTypes,
		DependencyIndexes: file_proto_omdb_proto_depIdxs,
		MessageInfos:      file_proto_omdb_proto_msgTypes,
	}.Build()
	File_proto_omdb_proto = out.File
	file_proto_omdb_proto_rawDesc = nil
	file_proto_omdb_proto_goTypes = nil
	file_proto_omdb_proto_depIdxs = nil
}
