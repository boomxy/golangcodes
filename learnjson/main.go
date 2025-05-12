package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// --- 结构体定义 ---
type Inner struct {
	Value int    `json:"value"`
	Label string `json:"label,omitempty"` // Label 带有 omitempty
}

type OuterStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Data Inner  `json:"data"` // 直接嵌套
}

type OuterPtr struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Data *Inner `json:"data"` // 指针嵌套
}

type OuterPtrOmit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Data *Inner `json:"data,omitempty"` // 指针嵌套 + omitempty
}

func marshalAndPrint(v interface{}, description string) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling %s: %v", description, err)
	}
	fmt.Printf("--- Marshalling %s ---\n%s\n", description, string(jsonData))
}

func unmarshalAndPrint(jsonStr string, v interface{}, description string) {
	err := json.Unmarshal([]byte(jsonStr), v)
	if err != nil {
		log.Fatalf("Error unmarshalling %s: %v", description, err)
	}
	fmt.Printf("--- Unmarshalling %s ---\nInput JSON:\n%s\nResulting Struct:\n%+v\n", description, jsonStr, v)
	// 特别检查指针类型的值
	if ptrStruct, ok := v.(*OuterPtr); ok {
		if ptrStruct.Data == nil {
			fmt.Println("  Inner Data Pointer: nil")
		} else {
			fmt.Printf("  Inner Data Pointer: %p, Value: %+v\n", ptrStruct.Data, *ptrStruct.Data)
		}
	}
	if ptrStructOmit, ok := v.(*OuterPtrOmit); ok {
		if ptrStructOmit.Data == nil {
			fmt.Println("  Inner Data Pointer: nil")
		} else {
			fmt.Printf("  Inner Data Pointer: %p, Value: %+v\n", ptrStructOmit.Data, *ptrStructOmit.Data)
		}
	}
	if normalStruct, ok := v.(*OuterStruct); ok {
		fmt.Printf("  Inner Data Value: %+v\n", normalStruct.Data)
	}
}

func main() {
	// --- 嵌套结构体 序列化 ---
	structPopulated := OuterStruct{ID: 1, Name: "Struct Populated", Data: Inner{Value: 100, Label: "Has Label"}}
	structZeroInner := OuterStruct{ID: 2, Name: "Struct Zero Inner", Data: Inner{}} // Inner 是零值
	marshalAndPrint(structPopulated, "OuterStruct (Populated Inner)")
	marshalAndPrint(structZeroInner, "OuterStruct (Zero Inner)") // Data 会序列化为 {"value":0}
	fmt.Println()
	fmt.Println("--------------------")
	fmt.Println()
	// --- 嵌套结构体指针 序列化 ---
	ptrPopulated := OuterPtr{ID: 11, Name: "Ptr Populated", Data: &Inner{Value: 110, Label: "Has Label"}}
	ptrNilInner := OuterPtr{ID: 12, Name: "Ptr Nil Inner", Data: nil}              // Data 是 nil
	ptrOmitNilInner := OuterPtrOmit{ID: 13, Name: "Ptr Omit Nil Inner", Data: nil} // Data 是 nil, 且有 omitempty
	marshalAndPrint(ptrPopulated, "OuterPtr (Populated Inner)")
	marshalAndPrint(ptrNilInner, "OuterPtr (Nil Inner)")         // Data 会序列化为 null
	marshalAndPrint(ptrOmitNilInner, "OuterPtrOmit (Nil Inner)") // Data 字段会被省略

	fmt.Println()
	fmt.Println("--------------------")
	fmt.Println()
	// --- 反序列化 JSON 示例 ---
	jsonWithData := `{"id": 101, "name": "JSON Has Data", "data": {"value": 200, "label": "From JSON"}}`
	jsonWithNull := `{"id": 102, "name": "JSON Null Data", "data": null}`
	jsonMissingData := `{"id": 103, "name": "JSON Missing Data"}`
	jsonWithZeroData := `{"id": 104, "name": "JSON Zero Data", "data": {"value": 0}}`

	// --- 反序列化到 嵌套结构体 ---
	var targetStruct OuterStruct
	unmarshalAndPrint(jsonWithData, &targetStruct, "OuterStruct <- JSON with Data")          // Data 被填充
	unmarshalAndPrint(jsonWithNull, &targetStruct, "OuterStruct <- JSON with Null")          // Data 变为 Inner{} 零值
	unmarshalAndPrint(jsonMissingData, &targetStruct, "OuterStruct <- JSON Missing Data")    // Data 变为 Inner{} 零值
	unmarshalAndPrint(jsonWithZeroData, &targetStruct, "OuterStruct <- JSON with Zero Data") // Data 变为 Inner{Value:0}

	fmt.Println()
	fmt.Println("--------------------")
	fmt.Println()

	// --- 反序列化到 嵌套结构体指针 ---
	var targetPtr OuterPtr
	unmarshalAndPrint(jsonWithData, &targetPtr, "OuterPtr <- JSON with Data")          // Data 指向新分配的 Inner
	unmarshalAndPrint(jsonWithNull, &targetPtr, "OuterPtr <- JSON with Null")          // Data 保持 nil
	unmarshalAndPrint(jsonMissingData, &targetPtr, "OuterPtr <- JSON Missing Data")    // Data 保持 nil
	unmarshalAndPrint(jsonWithZeroData, &targetPtr, "OuterPtr <- JSON with Zero Data") // Data 指向新分配的 Inner{Value:0}
}
