package client

// "parameterConfig": [
//   {
//     "description": "sha of the app to deploy",
//     "hasOptions": false,
//     "name": "version",
//     "options": [
//       {
//         "value": ""
//       }
//     ],
//     "required": true
//   },
//   {
//     "description": "commit message",
//     "name": "message"
//   },
//   {
//     "name": "committer_name"
//   },
//   {
//     "name": "committer_email"
//   }
// ],

type PipelineParameterOption struct {
	Value string `json:"value"`
}

type PipelineParameter struct {
	Description string                     `json:"description"`
	HasOptions  bool                       `json:"hasOptions"`
	Name        string                     `json:"name"`
	Options     *[]PipelineParameterOption `json:"options"`
	Required    bool                       `json:"required"`
}
