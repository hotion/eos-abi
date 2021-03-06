package eosabi

const abiDef = `
{
    "version": "eosio::abi/1.0",
    "types": [
        {
            "new_type_name": "key_name",
            "type": "string"
        },
        {
            "new_type_name": "key_type",
            "type": "string"
        },
        {
            "new_type_name": "action_name",
            "type": "name"
        },
        {
            "new_type_name": "table_name",
            "type": "name"
        }
    ],
    "structs": [
        {
            "name": "abi_def",
            "base": "",
            "fields": [
                {
                    "name": "version",
                    "type": "string"
                },
                {
                    "name": "types",
                    "type": "type[]"
                },
                {
                    "name": "structs",
                    "type": "struct[]"
                },
                {
                    "name": "actions",
                    "type": "action[]"
                },
                {
                    "name": "tables",
                    "type": "table[]"
                },
                {
                    "name": "ricardian_clauses",
                    "type": "ricardian_clause[]"
                },
                {
                    "name": "error_messages",
                    "type": "error_message[]"
                },
                {
                    "name": "abi_extensions",
                    "type": "extension[]"
                }
            ]
        },
        {
            "name": "type",
            "base": "",
            "fields": [
                {
                    "name": "new_type_name",
                    "type": "string"
                },
                {
                    "name": "type",
                    "type": "string"
                }
            ]
        },
        {
            "name": "struct",
            "base": "",
            "fields": [
                {
                    "name": "name",
                    "type": "string"
                },
                {
                    "name": "base",
                    "type": "string"
                },
                {
                    "name": "fields",
                    "type": "field[]"
                }
            ]
        },
        {
            "name": "field",
            "base": "",
            "fields": [
                {
                    "name": "name",
                    "type": "string"
                },
                {
                    "name": "type",
                    "type": "string"
                }
            ]
        },
        {
            "name": "action",
            "base": "",
            "fields": [
                {
                    "name": "name",
                    "type": "action_name"
                },
                {
                    "name": "type",
                    "type": "string"
                },
                {
                    "name": "ricardian_contract",
                    "type": "string"
                }
            ]
        },
        {
            "name": "table",
            "base": "",
            "fields": [
                {
                    "name": "name",
                    "type": "table_name"
                },
                {
                    "name": "index_type",
                    "type": "string"
                },
                {
                    "name": "key_names",
                    "type": "key_name[]"
                },
                {
                    "name": "key_types",
                    "type": "key_type[]"
                },
                {
                    "name": "type",
                    "type": "string"
                }
            ]
        },
        {
            "name": "ricardian_clause",
            "base": "",
            "fields": [
                {
                    "name": "id",
                    "type": "string"
                },
                {
                    "name": "body",
                    "type": "string"
                }
            ]
        },
        {
            "name": "error_message",
            "base": "",
            "fields": [
                {
                    "name": "error_code",
                    "type": "uint64"
                },
                {
                    "name": "error_msg",
                    "type": "string"
                }
            ]
        },
        {
            "name": "extension",
            "base": "",
            "fields": [
                {
                    "name": "id",
                    "type": "uint64"
                },
                {
                    "name": "ext",
                    "type": "bytes"
                }
            ]
        }
    ]
}
`
