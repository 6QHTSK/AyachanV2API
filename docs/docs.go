// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2/chart/metrics/bandori/{chartID}/{diffStr}": {
            "get": {
                "description": "该API可计算Bandori谱面的各项信息。\n\n### 标准谱面\n\n是指不需要玩家跨手/出张（左手跨过右手所在轨道）处理，且谱面的多压数不超过2的谱面，以及不存在重叠音符（两个音符在完全相同的轨道，并要求在完全相同的时间击打）\n\n\n### 基础信息\n\n是否非标准谱面、Note数、谱面时长、每秒平均谱面物件数（NPS）、是否有SP键、BPM情况、每秒平均击打数（HPS）、单位时间最大NPS（MaxScreenNPS）、谱面物件类型分布、谱面物件时间分布。\n\n### 标准谱面额外可计算的信息：\n\n左右手占比、左右手最大移动速度、单指最高每秒击打次数、单手粉键-音符平均间隔、单手音符-粉键平均间隔\n\n### 难度计算：\n\n基于统计回归的原理，通过拟合各个信息在各自难度所处于的位置，对比较的方式为每个信息进行难度标定。\n\n在当前版本，只会对基础信息部分的NPS、HPS、MaxScreenNPS三个维度进行难度回归计算，并对这三个信息的回归计算难度进行加权平均，给出拟合的**谱面总体难度**。\n\n额外可计算信息部分只会给出其相对于当前难度谱面的难度比较情况。例如，是否相对偏高/偏低。",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ChartMetrics"
                ],
                "summary": "计算Bandori谱面的各项信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "谱面ID",
                        "name": "chartID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "难度字符串，建议在[easy,normal,hard,expert,special]中选择",
                        "name": "diffStr",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description中所提及的各项信息",
                        "schema": {
                            "$ref": "#/definitions/model.ChartMetrics"
                        }
                    },
                    "400": {
                        "description": "传入Param错误",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误，包括找不到谱面等",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/chart/metrics/bestdori/{chartID}": {
            "get": {
                "description": "该API可计算Bandori谱面的各项信息。\n\n### 标准谱面\n\n是指不需要玩家跨手/出张（左手跨过右手所在轨道）处理，且谱面的多压数不超过2的谱面，以及不存在重叠音符（两个音符在完全相同的轨道，并要求在完全相同的时间击打）\n\n\n### 基础信息\n\n是否非标准谱面、Note数、谱面时长、每秒平均谱面物件数（NPS）、是否有SP键、BPM情况、每秒平均击打数（HPS）、单位时间最大NPS（MaxScreenNPS）、谱面物件类型分布、谱面物件时间分布。\n\n### 标准谱面额外可计算的信息：\n\n左右手占比、左右手最大移动速度、单指最高每秒击打次数、单手粉键-音符平均间隔、单手音符-粉键平均间隔\n\n### 难度计算：\n\n基于统计回归的原理，通过拟合各个信息在各自难度所处于的位置，对比较的方式为每个信息进行难度标定。\n\n在当前版本，只会对基础信息部分的NPS、HPS、MaxScreenNPS三个维度进行难度回归计算，并对这三个信息的回归计算难度进行加权平均，给出拟合的**谱面总体难度**。\n\n额外可计算信息部分只会给出其相对于当前难度谱面的难度比较情况。例如，是否相对偏高/偏低。",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ChartMetrics"
                ],
                "summary": "计算Bestdori谱面的各项信息，谱面的难度将会根据Bestdori上谱面声称的难度进行选择",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "谱面ID",
                        "name": "chartID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description中所提及的各项信息",
                        "schema": {
                            "$ref": "#/definitions/model.ChartMetrics"
                        }
                    },
                    "400": {
                        "description": "传入Param错误",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误，包括找不到谱面等",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/chart/metrics/custom/{diffStr}": {
            "post": {
                "description": "该API可计算Bandori谱面的各项信息。\n\n### 标准谱面\n\n是指不需要玩家跨手/出张（左手跨过右手所在轨道）处理，且谱面的多压数不超过2的谱面，以及不存在重叠音符（两个音符在完全相同的轨道，并要求在完全相同的时间击打）\n\n\n### 基础信息\n\n是否非标准谱面、Note数、谱面时长、每秒平均谱面物件数（NPS）、是否有SP键、BPM情况、每秒平均击打数（HPS）、单位时间最大NPS（MaxScreenNPS）、谱面物件类型分布、谱面物件时间分布。\n\n### 标准谱面额外可计算的信息：\n\n左右手占比、左右手最大移动速度、单指最高每秒击打次数、单手粉键-音符平均间隔、单手音符-粉键平均间隔\n\n### 难度计算：\n\n基于统计回归的原理，通过拟合各个信息在各自难度所处于的位置，对比较的方式为每个信息进行难度标定。\n\n在当前版本，只会对基础信息部分的NPS、HPS、MaxScreenNPS三个维度进行难度回归计算，并对这三个信息的回归计算难度进行加权平均，给出拟合的**谱面总体难度**。\n\n额外可计算信息部分只会给出其相对于当前难度谱面的难度比较情况。例如，是否相对偏高/偏低。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ChartMetrics"
                ],
                "summary": "计算上传谱面的各项信息，计算的各项信息请参考ChartMetricsFromBandori API文档所述。",
                "parameters": [
                    {
                        "type": "string",
                        "description": "难度字符串，建议在[easy,normal,hard,expert,special]中选择择",
                        "name": "diffStr",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "BestdoriV2谱面",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/bestdoriChart.BestdoriV2Note"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description中所提及的各项信息",
                        "schema": {
                            "$ref": "#/definitions/model.ChartMetrics"
                        }
                    },
                    "400": {
                        "description": "传入谱面/Param错误",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误，包括找不到谱面等",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/version": {
            "get": {
                "description": "根据内部信息得到API的版本",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Version"
                ],
                "summary": "获得API版本",
                "responses": {
                    "200": {
                        "description": "获得的API版本号",
                        "schema": {
                            "$ref": "#/definitions/controller.APIVersion"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "bestdoriChart.BestdoriV2BasicNote": {
            "type": "object",
            "properties": {
                "beat": {
                    "type": "number"
                },
                "flick": {
                    "type": "boolean"
                },
                "hidden": {
                    "type": "boolean"
                },
                "lane": {
                    "type": "number"
                }
            }
        },
        "bestdoriChart.BestdoriV2Note": {
            "type": "object",
            "properties": {
                "beat": {
                    "type": "number"
                },
                "bpm": {
                    "type": "number"
                },
                "connections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/bestdoriChart.BestdoriV2BasicNote"
                    }
                },
                "direction": {
                    "type": "string"
                },
                "flick": {
                    "type": "boolean"
                },
                "hidden": {
                    "type": "boolean"
                },
                "lane": {
                    "type": "number"
                },
                "type": {
                    "type": "string"
                },
                "width": {
                    "type": "integer"
                }
            }
        },
        "controller.APIVersion": {
            "type": "object",
            "properties": {
                "version": {
                    "type": "string"
                }
            }
        },
        "controller.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.ChartDifficultyExtend": {
            "type": "object",
            "properties": {
                "finger_max_hps": {
                    "type": "integer"
                },
                "flick_note_interval": {
                    "type": "integer"
                },
                "max_speed": {
                    "type": "integer"
                },
                "note_flick_interval": {
                    "type": "integer"
                }
            }
        },
        "model.ChartDifficultyStandard": {
            "type": "object",
            "properties": {
                "difficulty": {
                    "type": "number"
                },
                "max_screen_nps": {
                    "type": "number"
                },
                "total_hps": {
                    "type": "number"
                },
                "total_nps": {
                    "type": "number"
                }
            }
        },
        "model.ChartMetrics": {
            "type": "object",
            "properties": {
                "difficulty": {
                    "$ref": "#/definitions/model.ChartDifficultyStandard"
                },
                "difficulty_extend": {
                    "$ref": "#/definitions/model.ChartDifficultyExtend"
                },
                "metrics": {
                    "$ref": "#/definitions/model.ChartMetricsStandard"
                },
                "metrics_extend": {
                    "$ref": "#/definitions/model.ChartMetricsExtend"
                }
            }
        },
        "model.ChartMetricsExtend": {
            "type": "object",
            "properties": {
                "finger_max_hps": {
                    "type": "number"
                },
                "flick_note_interval": {
                    "type": "number"
                },
                "left_percent": {
                    "type": "number"
                },
                "max_speed": {
                    "type": "number"
                },
                "note_flick_interval": {
                    "type": "number"
                }
            }
        },
        "model.ChartMetricsStandard": {
            "type": "object",
            "properties": {
                "bpm_high": {
                    "type": "number"
                },
                "bpm_low": {
                    "type": "number"
                },
                "distribution": {
                    "$ref": "#/definitions/model.Distribution"
                },
                "irregular": {
                    "description": "存在多压/交叉（出张）0 失败 1 标准 2 非标准",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.RegularType"
                        }
                    ]
                },
                "irregular_info": {
                    "description": "无法分析的第一个错误情况",
                    "type": "string"
                },
                "main_bpm": {
                    "type": "number"
                },
                "max_screen_nps": {
                    "type": "number"
                },
                "noteCount": {
                    "$ref": "#/definitions/model.NoteCount"
                },
                "sp_rhythm": {
                    "type": "boolean"
                },
                "total_hit_note": {
                    "type": "integer"
                },
                "total_hps": {
                    "type": "number"
                },
                "total_note": {
                    "type": "integer"
                },
                "total_nps": {
                    "type": "number"
                },
                "total_time": {
                    "type": "number"
                }
            }
        },
        "model.Distribution": {
            "type": "object",
            "properties": {
                "hit": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "note": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.NoteCount": {
            "type": "object",
            "properties": {
                "direction_left": {
                    "type": "integer"
                },
                "direction_right": {
                    "type": "integer"
                },
                "flick": {
                    "type": "integer"
                },
                "single": {
                    "type": "integer"
                },
                "slide_end": {
                    "type": "integer"
                },
                "slide_flick": {
                    "type": "integer"
                },
                "slide_hidden": {
                    "type": "integer"
                },
                "slide_start": {
                    "type": "integer"
                },
                "slide_tick": {
                    "type": "integer"
                }
            }
        },
        "model.RegularType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "RegularTypeUnknown",
                "RegularTypeRegular",
                "RegularTypeIrregular"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.2",
	Host:             "",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "Ayachan Bandori谱面难度分析器",
	Description:      "# Ayachan Bandori谱面难度分析器\n\n[![codebeat badge](https://codebeat.co/badges/3482bd1e-45d7-4e83-af70-3f1ccb874fcd)](https://codebeat.co/projects/github-com-6qhtsk-ayachan-development)\n[![Go Report Card](https://goreportcard.com/badge/github.com/6QHTSK/ayachan)](https://goreportcard.com/report/github.com/6QHTSK/ayachan)\n![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/6QHTSK/ayachan)\n![GitHub](https://img.shields.io/github/license/6QHTSK/ayachan)\n![Libraries.io dependency status for GitHub repo](https://img.shields.io/librariesio/github/6QHTSK/ayachan)\n\n可对Bandori谱面进行特征提取，并拟合难度\n",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
