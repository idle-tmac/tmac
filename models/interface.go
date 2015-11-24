package models

type Singleton interface { 
    GetInstance() (interface{})
} 
