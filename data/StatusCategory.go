package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -pkg jiradata -overwrite ../schemas/SearchResults.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// StatusCategory defined from schema:
// {
//   "title": "Status Category",
//   "type": "object",
//   "properties": {
//     "colorName": {
//       "title": "colorName",
//       "type": "string"
//     },
//     "id": {
//       "title": "id",
//       "type": "integer"
//     },
//     "key": {
//       "title": "key",
//       "type": "string"
//     },
//     "name": {
//       "title": "name",
//       "type": "string"
//     },
//     "self": {
//       "title": "self",
//       "type": "string"
//     }
//   }
// }
type StatusCategory struct {
	ColorName string `json:"colorName,omitempty" yaml:"colorName,omitempty"`
	ID        int    `json:"id,omitempty" yaml:"id,omitempty"`
	Key       string `json:"key,omitempty" yaml:"key,omitempty"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Self      string `json:"self,omitempty" yaml:"self,omitempty"`
}
