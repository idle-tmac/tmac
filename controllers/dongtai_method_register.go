package controllers 
import (
	_"fmt"
        _"tmac/common"
        _"tmac/models"
        _"tmac/lib"
        _"encoding/json"
        _"strconv"
        _"github.com/astaxie/beego/cache"
        _"github.com/astaxie/beego/cache/redis" 
        _"log" 
        _"time"
)


//dongtai recommend function register
var dongtaiFuncs = map[string]interface{} {
	"recommend_cells":ReqRecommendCells,
}
