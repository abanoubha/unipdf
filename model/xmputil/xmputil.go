//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package xmputil provides abstraction used by the pdf document XMP Metadata.
package xmputil ;import (_dg "errors";_ag "fmt";_fbe "github.com/trimmer-io/go-xmp/models/pdf";_ed "github.com/trimmer-io/go-xmp/models/xmp_base";_e "github.com/trimmer-io/go-xmp/models/xmp_mm";_fb "github.com/trimmer-io/go-xmp/xmp";_a "github.com/unidoc/unipdf/v4/core";
_df "github.com/unidoc/unipdf/v4/internal/timeutils";_ca "github.com/unidoc/unipdf/v4/internal/uuid";_g "github.com/unidoc/unipdf/v4/model/xmputil/pdfaextension";_cc "github.com/unidoc/unipdf/v4/model/xmputil/pdfaid";_d "strconv";_c "time";);

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_deg *Document )SetPdfAID (part int ,conformance string )error {_gg ,_abg :=_cc .MakeModel (_deg ._b );if _abg !=nil {return _abg ;};_gg .Part =part ;_gg .Conformance =conformance ;if _ce :=_gg .SyncToXMP (_deg ._b );_ce !=nil {return _ce ;};return nil ;
};

// GetPdfAID gets the pdfaid xmp metadata model.
func (_ggf *Document )GetPdfAID ()(*PdfAID ,bool ){_bge ,_bb :=_ggf ._b .FindModel (_cc .Namespace ).(*_cc .Model );if !_bb {return nil ,false ;};return &PdfAID {Part :_bge .Part ,Conformance :_bge .Conformance },true ;};

// SetPdfInfo sets the pdf info into selected document.
func (_ga *Document )SetPdfInfo (options *PdfInfoOptions )error {if options ==nil {return _dg .New ("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};_fd ,_cb :=_fbe .MakeModel (_ga ._b );
if _cb !=nil {return _cb ;};if options .Overwrite {*_fd =_fbe .PDFInfo {};};if options .InfoDict !=nil {_aeb ,_ff :=_a .GetDict (options .InfoDict );if !_ff {return _ag .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",options .InfoDict );
};var _egf *_a .PdfObjectString ;for _ ,_gda :=range _aeb .Keys (){switch _gda {case "\u0054\u0069\u0074l\u0065":_egf ,_ff =_a .GetString (_aeb .Get ("\u0054\u0069\u0074l\u0065"));if _ff {_fd .Title =_fb .NewAltString (_egf );};case "\u0041\u0075\u0074\u0068\u006f\u0072":_egf ,_ff =_a .GetString (_aeb .Get ("\u0041\u0075\u0074\u0068\u006f\u0072"));
if _ff {_fd .Author =_fb .NewStringList (_egf .String ());};case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":_egf ,_ff =_a .GetString (_aeb .Get ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"));if _ff {_fd .Keywords =_egf .String ();};case "\u0043r\u0065\u0061\u0074\u006f\u0072":_egf ,_ff =_a .GetString (_aeb .Get ("\u0043r\u0065\u0061\u0074\u006f\u0072"));
if _ff {_fd .Creator =_fb .AgentName (_egf .String ());};case "\u0053u\u0062\u006a\u0065\u0063\u0074":_egf ,_ff =_a .GetString (_aeb .Get ("\u0053u\u0062\u006a\u0065\u0063\u0074"));if _ff {_fd .Subject =_fb .NewAltString (_egf .String ());};case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":_egf ,_ff =_a .GetString (_aeb .Get ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"));
if _ff {_fd .Producer =_fb .AgentName (_egf .String ());};case "\u0054r\u0061\u0070\u0070\u0065\u0064":_dd ,_ge :=_a .GetName (_aeb .Get ("\u0054r\u0061\u0070\u0070\u0065\u0064"));if _ge {switch _dd .String (){case "\u0054\u0072\u0075\u0065":_fd .Trapped =true ;
case "\u0046\u0061\u006cs\u0065":_fd .Trapped =false ;default:_fd .Trapped =true ;};};case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":if _egg ,_de :=_a .GetString (_aeb .Get ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"));
_de &&_egg .String ()!=""{_fde ,_gcd :=_df .ParsePdfTime (_egg .String ());if _gcd !=nil {return _ag .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077",_gcd );
};_fd .CreationDate =_fb .NewDate (_fde );};case "\u004do\u0064\u0044\u0061\u0074\u0065":if _be ,_gdaa :=_a .GetString (_aeb .Get ("\u004do\u0064\u0044\u0061\u0074\u0065"));_gdaa &&_be .String ()!=""{_ac ,_dfc :=_df .ParsePdfTime (_be .String ());if _dfc !=nil {return _ag .Errorf ("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w",_dfc );
};_fd .ModifyDate =_fb .NewDate (_ac );};};};};if options .PdfVersion !=""{_fd .PDFVersion =options .PdfVersion ;};if options .Marked {_fd .Marked =_fb .Bool (options .Marked );};if options .Copyright !=""{_fd .Copyright =options .Copyright ;};if _cb =_fd .SyncToXMP (_ga ._b );
_cb !=nil {return _cb ;};return nil ;};

// Marshal the document into xml byte stream.
func (_eda *Document )Marshal ()([]byte ,error ){if _eda ._b .IsDirty (){if _gd :=_eda ._b .SyncModels ();_gd !=nil {return nil ,_gd ;};};return _fb .Marshal (_eda ._b );};

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_fc *Document )MarshalIndent (prefix ,indent string )([]byte ,error ){if _fc ._b .IsDirty (){if _gc :=_fc ._b .SyncModels ();_gc !=nil {return nil ,_gc ;};};return _fb .MarshalIndent (_fc ._b ,prefix ,indent );};

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct{VersionID string ;ModifyDate _c .Time ;Comments string ;Modifier string ;};

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_ccc *Document )SetMediaManagement (options *MediaManagementOptions )error {_fe ,_acd :=_e .MakeModel (_ccc ._b );if _acd !=nil {return _acd ;};if options ==nil {options =new (MediaManagementOptions );};_afc :=_e .ResourceRef {};switch {case options .DocumentID !="":_fe .DocumentID =_fb .GUID (options .DocumentID );
case options .NewDocumentID ||_fe .DocumentID .IsZero ():if !_fe .DocumentID .IsZero (){_afc .DocumentID =_fe .DocumentID ;};_bg ,_gb :=_ca .NewUUID ();if _gb !=nil {return _gb ;};_fe .DocumentID =_fb .GUID (_bg .String ());};if !_fe .InstanceID .IsZero (){_afc .InstanceID =_fe .InstanceID ;
};_fe .InstanceID =_fb .GUID (options .InstanceID );if _fe .InstanceID ==""{_gf ,_ef :=_ca .NewUUID ();if _ef !=nil {return _ef ;};_fe .InstanceID =_fb .GUID (_gf .String ());};if !_afc .IsZero (){_fe .DerivedFrom =&_afc ;};_agg :=options .VersionID ;if _fe .VersionID !=""{_cae ,_ffe :=_d .Atoi (_fe .VersionID );
if _ffe !=nil {_agg =_d .Itoa (len (_fe .Versions )+1);}else {_agg =_d .Itoa (_cae +1);};};if _agg ==""{_agg ="\u0031";};_fe .VersionID =_agg ;if _acd =_fe .SyncToXMP (_ccc ._b );_acd !=nil {return _acd ;};return nil ;};

// MediaManagement are the values from the document media management metadata.
type MediaManagement struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
OriginalDocumentID GUID ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
DocumentID GUID ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
InstanceID GUID ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
DerivedFrom *MediaManagementDerivedFrom ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
VersionID string ;

// Versions is the history of the document versions along with the comments, timestamps and issuers.
Versions []MediaManagementVersion ;};

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_eg *Document )GetGoXmpDocument ()*_fb .Document {return _eg ._b };

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_ab *Document )GetMediaManagement ()(*MediaManagement ,bool ){_eeb :=_e .FindModel (_ab ._b );if _eeb ==nil {return nil ,false ;};_ea :=make ([]MediaManagementVersion ,len (_eeb .Versions ));for _eaa ,_gag :=range _eeb .Versions {_ea [_eaa ]=MediaManagementVersion {VersionID :_gag .Version ,ModifyDate :_gag .ModifyDate .Value (),Comments :_gag .Comments ,Modifier :_gag .Modifier };
};_dfcc :=&MediaManagement {OriginalDocumentID :GUID (_eeb .OriginalDocumentID .Value ()),DocumentID :GUID (_eeb .DocumentID .Value ()),InstanceID :GUID (_eeb .InstanceID .Value ()),VersionID :_eeb .VersionID ,Versions :_ea };if _eeb .DerivedFrom !=nil {_dfcc .DerivedFrom =&MediaManagementDerivedFrom {OriginalDocumentID :GUID (_eeb .DerivedFrom .OriginalDocumentID ),DocumentID :GUID (_eeb .DerivedFrom .DocumentID ),InstanceID :GUID (_eeb .DerivedFrom .InstanceID ),VersionID :_eeb .DerivedFrom .VersionID };
};return _dfcc ,true ;};

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_bf *Document )GetPdfaExtensionSchemas ()([]_g .Schema ,error ){_bdc :=_bf ._b .FindModel (_g .Namespace );if _bdc ==nil {return nil ,nil ;};_db ,_ede :=_bdc .(*_g .Model );if !_ede {return nil ,_ag .Errorf ("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054",_bdc );
};return _db .Schemas ,nil ;};

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct{OriginalDocumentID GUID ;DocumentID GUID ;InstanceID GUID ;VersionID string ;};

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_bd *Document )SetPdfAExtension ()error {_fbg ,_ae :=_g .MakeModel (_bd ._b );if _ae !=nil {return _ae ;};if _ae =_g .FillModel (_bd ._b ,_fbg );_ae !=nil {return _ae ;};if _ae =_fbg .SyncToXMP (_bd ._b );_ae !=nil {return _ae ;};return nil ;};

// GUID is a string representing a globally unique identifier.
type GUID string ;

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct{Part int ;Conformance string ;};

// PdfInfo is the xmp document pdf info.
type PdfInfo struct{InfoDict _a .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;};

// MediaManagementOptions are the options for the Media management xmp metadata.
type MediaManagementOptions struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
// By default, this value is generated.
OriginalDocumentID string ;

// NewDocumentID is a flag which generates a new Document identifier while setting media management.
// This value should be set to true only if the document is stored and saved as new document.
// Otherwise, if the document is modified and overwrites previous file, it should be set to false.
NewDocumentID bool ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
// By default, this value is generated if NewDocumentID is true or previous doesn't exist.
DocumentID string ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
// By default, this value is generated.
InstanceID string ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
// By default, the derived from structure is filled from previous XMP metadata (if exists).
DerivedFrom string ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
// By default, this values is incremented or set to the next version number.
VersionID string ;

// ModifyComment is a comment to given modification
ModifyComment string ;

// ModifyDate is a custom modification date for the versions.
// By default, this would be set to time.Now().
ModifyDate _c .Time ;

// Modifier is a person who did the modification.
Modifier string ;};

// GetPdfInfo gets the document pdf info.
func (_ffg *Document )GetPdfInfo ()(*PdfInfo ,bool ){_agf :=PdfInfo {};var _caf *_a .PdfObjectDictionary ;_aga :=func (_dfe string ,_dga _a .PdfObject ){if _caf ==nil {_caf =_a .MakeDict ();};_caf .Set (_a .PdfObjectName (_dfe ),_dga );};_dc ,_afg :=_ffg ._b .FindModel (_fbe .NsPDF ).(*_fbe .PDFInfo );
if !_afg {_beb ,_da :=_ffg ._b .FindModel (_ed .NsXmp ).(*_ed .XmpBase );if !_da {return nil ,false ;};if _beb .CreatorTool !=""{_aga ("\u0043r\u0065\u0061\u0074\u006f\u0072",_a .MakeString (string (_beb .CreatorTool )));};if !_beb .CreateDate .IsZero (){_aga ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_a .MakeString (_df .FormatPdfTime (_beb .CreateDate .Value ())));
};if !_beb .ModifyDate .IsZero (){_aga ("\u004do\u0064\u0044\u0061\u0074\u0065",_a .MakeString (_df .FormatPdfTime (_beb .ModifyDate .Value ())));};_agf .InfoDict =_caf ;return &_agf ,true ;};_agf .Copyright =_dc .Copyright ;_agf .PdfVersion =_dc .PDFVersion ;
_agf .Marked =bool (_dc .Marked );if len (_dc .Title )> 0{_aga ("\u0054\u0069\u0074l\u0065",_a .MakeString (_dc .Title .Default ()));};if len (_dc .Author )> 0{_aga ("\u0041\u0075\u0074\u0068\u006f\u0072",_a .MakeString (_dc .Author [0]));};if _dc .Keywords !=""{_aga ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",_a .MakeString (_dc .Keywords ));
};if len (_dc .Subject )> 0{_aga ("\u0053u\u0062\u006a\u0065\u0063\u0074",_a .MakeString (_dc .Subject .Default ()));};if _dc .Creator !=""{_aga ("\u0043r\u0065\u0061\u0074\u006f\u0072",_a .MakeString (string (_dc .Creator )));};if _dc .Producer !=""{_aga ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",_a .MakeString (string (_dc .Producer )));
};if _dc .Trapped {_aga ("\u0054r\u0061\u0070\u0070\u0065\u0064",_a .MakeName ("\u0054\u0072\u0075\u0065"));};if !_dc .CreationDate .IsZero (){_aga ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_a .MakeString (_df .FormatPdfTime (_dc .CreationDate .Value ())));
};if !_dc .ModifyDate .IsZero (){_aga ("\u004do\u0064\u0044\u0061\u0074\u0065",_a .MakeString (_df .FormatPdfTime (_dc .ModifyDate .Value ())));};_agf .InfoDict =_caf ;return &_agf ,true ;};

// NewDocument creates a new document without any previous xmp information.
func NewDocument ()*Document {_af :=_fb .NewDocument ();return &Document {_b :_af }};

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument (stream []byte )(*Document ,error ){_ee :=_fb .NewDocument ();if _ec :=_fb .Unmarshal (stream ,_ee );_ec !=nil {return nil ,_ec ;};return &Document {_b :_ee },nil ;};

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{_b *_fb .Document };

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct{InfoDict _a .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;

// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
Overwrite bool ;};