# binance-kline-provider

### GetPreviousKlines

* GET /api/kline/`symbol`/`interval`/previous

    |   参数名  	    |  是否必选 	        |   默认值               |  描述	|
    |-------------	|  --------------	|----------------------	|------------------------	|
    | endTime       |       否      	| 2147483647000         | 时间戳(毫秒)         	    |
    | limit      	|       否      	| 1000             	    | 向前获取数量(不得超过1000)  	|

* example: 
    - 请求 `GET /api/kline/BTCUSDT/1m/previous?endTime=1503360000000&&limit=5`
    - 返回
    ```json
    [{
        "open_time": 1503360000000,
        "open_price": "4016",
        "close_price": "4016",
        "high_price": "4016",
        "low_price": "4016"
    }, {
        "open_time": 1503359940000,
        "open_price": "4016",
        "close_price": "4016",
        "high_price": "4016",
        "low_price": "4016"
    }, {
        "open_time": 1503359880000,
        "open_price": "4016",
        "close_price": "4016",
        "high_price": "4016",
        "low_price": "4016"
    }, {
        "open_time": 1503359820000,
        "open_price": "4010",
        "close_price": "4016",
        "high_price": "4010",
        "low_price": "4016"
    }, {
        "open_time": 1503359760000,
        "open_price": "4010",
        "close_price": "4010",
        "high_price": "4010",
        "low_price": "4010"
    }]
    ```

### GetNextKlines 

* GET /api/kline/`symbol`/`interval`/next

    |   参数名  	    |  是否必选 	        |   默认值               |  描述	|
    |-------------	|  --------------	|----------------------	|------------------------	|
    | fromTime      |       否      	| 0                     | 时间戳(毫秒)         	    |
    | limit      	|       否      	| 1000             	    | 向后获取数量(不得超过1000)  	|

* example: 
    - 请求 `GET /api/kline/BTCUSDT/1m/next?fromTime=1503360000000&&limit=5`
    - 返回
    ```json
    [{
        "open_time": 1503360000000,
        "open_price": "4016",
        "close_price": "4016",
        "high_price": "4016",
        "low_price": "4016"
    }, {
        "open_time": 1503360060000,
        "open_price": "4016",
        "close_price": "4016",
        "high_price": "4016",
        "low_price": "4016"
    }, {
        "open_time": 1503360120000,
        "open_price": "4016",
        "close_price": "4016",
        "high_price": "4016",
        "low_price": "4016"
    }, {
        "open_time": 1503360180000,
        "open_price": "4016",
        "close_price": "4016",
        "high_price": "4016",
        "low_price": "4016"
    }, {
        "open_time": 1503360240000,
        "open_price": "4016",
        "close_price": "4016",
        "high_price": "4016",
        "low_price": "4016"
    }]
    ```

