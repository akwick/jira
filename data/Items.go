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

// Items defined from schema:
// {
//   "title": "items",
//   "type": "array",
//   "items": {
//     "title": "Change Item",
//     "type": "object",
//     "properties": {
//       "field": {
//         "title": "field",
//         "type": "string"
//       },
//       "fieldtype": {
//         "title": "fieldtype",
//         "type": "string"
//       },
//       "from": {
//         "title": "from",
//         "type": "string"
//       },
//       "fromString": {
//         "title": "fromString",
//         "type": "string"
//       },
//       "to": {
//         "title": "to",
//         "type": "string"
//       },
//       "toString": {
//         "title": "toString",
//         "type": "string"
//       }
//     }
//   }
// }
type Items []*ChangeItem
