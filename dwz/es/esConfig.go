package main

import elastic "gopkg.in/olivere/elastic.v5"

var esClient *elastic.Client

const indexName string = "dwzindex"
const indexBody string = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"_default_": {
			"_all": {
				"enabled": true
			}
		},
		"urlinfo":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				}
			}
		},
		"comment":{
			"_parent": {
				"type":	"urlinfo"
			}
		},
		"order":{
			"properties":{
				"article":{
					"type":"text"
				},
				"manufacturer":{
					"type":"keyword"
				},
				"price":{
					"type":"float"
				},
				"time":{
					"type":"date",
					"format": "YYYY-MM-dd"
				}
			}
		},
		"doctype":{
			"properties":{
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				}
			}
		},
		"queries":{
			"properties": {
				"query": {
					"type":	"percolator"
				}
			}
		},
		"tweet-nosource":{
			"_source": {
				"enabled": false
			},
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion",
					"contexts":[
						{
							"name":"user_name",
							"type":"category"
						}
					]
				}
			}
		}
	}
}
`
