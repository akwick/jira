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

// Transition defined from schema:
// {
//   "title": "Transition",
//   "type": "object",
//   "properties": {
//     "expand": {
//       "title": "expand",
//       "type": "string"
//     },
//     "fields": {
//       "title": "fields",
//       "type": "object",
//       "patternProperties": {
//         ".+": {
//           "title": "Field Meta",
//           "type": "object",
//           "properties": {
//             "allowedValues": {
//               "type": "array",
//               "items": {}
//             },
//             "autoCompleteUrl": {
//               "type": "string"
//             },
//             "hasDefaultValue": {
//               "type": "boolean"
//             },
//             "key": {
//               "type": "string"
//             },
//             "name": {
//               "type": "string"
//             },
//             "operations": {
//               "type": "array",
//               "items": {
//                 "type": "string"
//               }
//             },
//             "required": {
//               "type": "boolean"
//             },
//             "schema": {
//               "$ref": "#/definitions/json-type"
//             }
//           }
//         }
//       }
//     },
//     "hasScreen": {
//       "title": "hasScreen",
//       "type": "boolean"
//     },
//     "id": {
//       "title": "id",
//       "type": "string"
//     },
//     "name": {
//       "title": "name",
//       "type": "string"
//     },
//     "to": {
//       "title": "Status",
//       "type": "object",
//       "properties": {
//         "description": {
//           "title": "description",
//           "type": "string"
//         },
//         "iconUrl": {
//           "title": "iconUrl",
//           "type": "string"
//         },
//         "id": {
//           "title": "id",
//           "type": "string"
//         },
//         "name": {
//           "title": "name",
//           "type": "string"
//         },
//         "self": {
//           "title": "self",
//           "type": "string"
//         },
//         "statusCategory": {
//           "title": "Status Category",
//           "type": "object",
//           "properties": {
//             "colorName": {
//               "title": "colorName",
//               "type": "string"
//             },
//             "id": {
//               "title": "id",
//               "type": "integer"
//             },
//             "key": {
//               "title": "key",
//               "type": "string"
//             },
//             "name": {
//               "title": "name",
//               "type": "string"
//             },
//             "self": {
//               "title": "self",
//               "type": "string"
//             }
//           }
//         },
//         "statusColor": {
//           "title": "statusColor",
//           "type": "string"
//         }
//       }
//     }
//   }
// }
type Transition struct {
	Expand    string       `json:"expand,omitempty" yaml:"expand,omitempty"`
	Fields    FieldMetaMap `json:"fields,omitempty" yaml:"fields,omitempty"`
	HasScreen bool         `json:"hasScreen,omitempty" yaml:"hasScreen,omitempty"`
	ID        string       `json:"id,omitempty" yaml:"id,omitempty"`
	Name      string       `json:"name,omitempty" yaml:"name,omitempty"`
	To        *Status      `json:"to,omitempty" yaml:"to,omitempty"`
}
