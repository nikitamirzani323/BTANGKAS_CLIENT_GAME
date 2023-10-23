<script>
    import dayjs from "dayjs";
    import utc from "dayjs/plugin/utc";
    import timezone from "dayjs/plugin/timezone";
    
    dayjs.extend(utc);
    dayjs.extend(timezone);
    export let path_api = "";
    export let token_browser = "";
    export let client_company = "";
    export let client_timezone = "";
    export let client_ipaddress = "";
    export let client_name = "";
    export let client_username = "";
    export let client_credit = 0;
    export let client_listbet = [];
  
   
    let clockmachine = "";
    let clockmachine_data = "";
    let credit = client_credit;
    let temp_credit = credit;
    let list_min_bet = [100,500,1000,1500,5000,10000,25000]
    let idtransaksi = ""
    let tipe_click = ""
    let min_bet = "0"
    let totalbet = 0
    let sound = 0;
    let flag_minimalbet = false;
    let path_card = "CARD/WHITE/";
    let list_point_db = [];
    if(client_listbet != null){
      for(let i=0;i<client_listbet.length;i++){
        if(i==0){
          min_bet = client_listbet[i].lisbet_minbet
          for(let j=0;j<client_listbet[i].lisbet_config.length;j++){
            list_point_db = [
                ...list_point_db,
                {
                  lispoin_id: client_listbet[i].lisbet_config[j].lispoin_id,
                  lispoin_nmpoin: client_listbet[i].lisbet_config[j].lispoin_nmpoin,
                  lispoin_poin: client_listbet[i].lisbet_config[j].lispoin_poin,
                },
            ];
          }
        }
        
      }
    }
    // console.log(list_point_db)
    let spin = [
      new Audio("/sounds/spin.mp3"),
    ];
    let win = [
      new Audio("/sounds/win.mp3"),
      new Audio("/sounds/win.mp3"),
    ];
    // let list_point22 = [
    //   {id:"1",code:"RF",name:"Royal Flush",poin:500},
    //   {id:"2",code:"5K",name:"5 Of A Kind",poin:200},
    //   {id:"3",code:"SF",name:"Straight Flush",poin:120},
    //   {id:"4",code:"4K",name:"4 Of A Kind",poin:50},
    //   {id:"5",code:"FH",name:"Full House",poin:7},
    //   {id:"6",code:"FL",name:"Flush",poin:5},
    //   {id:"7",code:"ST",name:"Straight",poin:3},
    //   {id:"8",code:"3K",name:"3 Of A Kind",poin:2},
    //   {id:"9",code:"2P",name:"2 Pair (10 PAIR)",poin:1},
    //   {id:"10",code:"AP",name:"Ace Pair",poin:1},
    // ]
    let list_point = [
      {id:"1",code:"RF",name:"Royal Flush",poin:0},
      {id:"2",code:"5K",name:"5 Of A Kind",poin:0},
      {id:"3",code:"SF",name:"Straight Flush",poin:0},
      {id:"4",code:"4K",name:"4 Of A Kind",poin:0},
      {id:"5",code:"FH",name:"Full House",poin:0},
      {id:"6",code:"FL",name:"Flush",poin:0},
      {id:"7",code:"ST",name:"Straight",poin:0},
      {id:"8",code:"3K",name:"3 Of A Kind",poin:0},
      {id:"9",code:"2P",name:"2 Pair (10 PAIR)",poin:0},
      {id:"10",code:"AP",name:"Ace Pair",poin:0},
    ]
    const card_result_data = [
      {id:"2_diamond",val:"2",val_display:2,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_2.png"},
      {id:"3_diamond",val:"3",val_display:3,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_3.png"},
      {id:"4_diamond",val:"4",val_display:4,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_4.png"},
      {id:"5_diamond",val:"5",val_display:5,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_5.png"},
      {id:"6_diamond",val:"6",val_display:6,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_6.png"},
      {id:"7_diamond",val:"7",val_display:7,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_7.png"},
      {id:"8_diamond",val:"8",val_display:8,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_8.png"},
      {id:"9_diamond",val:"9",val_display:9,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_9.png"},
      {id:"10_diamond",val:"10",val_display:10,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_10.png"},
      {id:"j_diamond",val:"J",val_display:11,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_J.png"},
      {id:"q_diamond",val:"Q",val_display:12,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_Q.png"},
      {id:"k_diamond",val:"K",val_display:13,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_K.png"},
      {id:"as_diamond",val:"AS",val_display:14,code_card:"D",img:"./CARD/WHITE/CARD_RED_DIAMOND_AS.png"},
      {id:"2_heart",val:"2",val_display:2,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_2.png"},
      {id:"3_heart",val:"3",val_display:3,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_3.png"},
      {id:"4_heart",val:"4",val_display:4,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_4.png"},
      {id:"5_heart",val:"5",val_display:5,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_5.png"},
      {id:"6_heart",val:"6",val_display:6,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_6.png"},
      {id:"7_heart",val:"7",val_display:7,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_7.png"},
      {id:"8_heart",val:"8",val_display:8,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_8.png"},
      {id:"9_heart",val:"9",val_display:9,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_9.png"},
      {id:"10_heart",val:"10",val_display:10,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_10.png"},
      {id:"j_heart",val:"J",val_display:11,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_J.png"},
      {id:"q_heart",val:"Q",val_display:12,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_Q.png"},
      {id:"k_heart",val:"K",val_display:13,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_K.png"},
      {id:"as_heart",val:"AS",val_display:14,code_card:"H",img:"./CARD/WHITE/CARD_RED_LOVE_AS.png"},
      {id:"2_club",val:"2",val_display:2,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_2.png"},
      {id:"3_club",val:"3",val_display:3,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_3.png"},
      {id:"4_club",val:"4",val_display:4,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_4.png"},
      {id:"5_club",val:"5",val_display:5,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_5.png"},
      {id:"6_club",val:"6",val_display:6,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_6.png"},
      {id:"7_club",val:"7",val_display:7,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_7.png"},
      {id:"8_club",val:"8",val_display:8,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_8.png"},
      {id:"9_club",val:"9",val_display:9,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_9.png"},
      {id:"10_club",val:"10",val_display:10,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_10.png"},
      {id:"j_club",val:"J",val_display:11,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_J.png"},
      {id:"q_club",val:"Q",val_display:12,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_Q.png"},
      {id:"k_club",val:"K",val_display:13,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_K.png"},
      {id:"as_club",val:"AS",val_display:14,code_card:"C",img:"./CARD/WHITE/CARD_BLACK_KRITING_AS.png"},
      {id:"2_spade",val:"2",val_display:2,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_2.png"},
      {id:"3_spade",val:"3",val_display:3,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_3.png"},
      {id:"4_spade",val:"4",val_display:4,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_4.png"},
      {id:"5_spade",val:"5",val_display:5,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_5.png"},
      {id:"6_spade",val:"6",val_display:6,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_6.png"},
      {id:"7_spade",val:"7",val_display:7,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_7.png"},
      {id:"8_spade",val:"8",val_display:8,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_8.png"},
      {id:"9_spade",val:"9",val_display:9,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_9.png"},
      {id:"10_spade",val:"10",val_display:10,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_10.png"},
      {id:"j_spade",val:"J",val_display:11,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_J.png"},
      {id:"q_spade",val:"Q",val_display:12,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_Q.png"},
      {id:"k_spade",val:"K",val_display:13,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_K.png"},
      {id:"as_spade",val:"AS",val_display:14,code_card:"S",img:"./CARD/WHITE/CARD_BLACK_DAUN_AS.png"},
      {id:"jk_black",val:"JK",val_display:1,code_card:"JK",img:"./CARD/WHITE/CARD_JOKER_BLACK.png"},
      {id:"jk_red",val:"JK",val_display:1,code_card:"JK",img:"./CARD/WHITE/CARD_JOKER_RED.png"},
    ]
    function updateClock() {
      let endtime = dayjs().tz(client_timezone).format("DD MMM YYYY | HH:mm");
      let endtime_data = dayjs().tz(client_timezone).format("YYYY-MM-DD HH:mm:ss");
      clockmachine = endtime;
      clockmachine_data = endtime_data;
    }
    const pattern_stright_1 = [2,3,4,5,6]
    const pattern_stright_2 = [3,4,5,6,7]
    const pattern_stright_3 = [4,5,6,7,8]
    const pattern_stright_4 = [5,6,7,8,9]
    const pattern_stright_5 = [6,7,8,9,10]
    const pattern_stright_6 = [7,8,9,10,11]
    const pattern_stright_7 = [8,9,10,11,12]
    const pattern_stright_8 = [9,10,11,12,13]
    const pattern_stright_9 = [10,11,12,13,14]
    const pattern_stright_10 = [14,2,3,4,5]

    const pattern_4_1 = [2,3,4,5]
    const pattern_4_2 = [3,4,5,6]
    const pattern_4_3 = [4,5,6,7]
    const pattern_4_4 = [5,6,7,8]
    const pattern_4_5 = [6,7,8,9]
    const pattern_4_6 = [7,8,9,10]
    const pattern_4_7 = [8,9,10,11]
    const pattern_4_8 = [9,10,11,12]
    const pattern_4_9 = [10,11,12,13]
    const pattern_4_10 = [11,12,13,14]
    const pattern_4_11 = [14,2,3,4]

    const pattern_3_1 = [2,3,4]
    const pattern_3_2 = [3,4,5]
    const pattern_3_3 = [4,5,6]
    const pattern_3_4 = [5,6,7]
    const pattern_3_5 = [6,7,8]
    const pattern_3_6 = [7,8,9]
    const pattern_3_7 = [8,9,10]
    const pattern_3_8 = [9,10,11]
    const pattern_3_9 = [10,11,12]
    const pattern_3_10 = [11,12,13]
    const pattern_3_11 = [12,13,14]
    const pattern_3_12 = [14,2,3]
    
    let list_invoice = []
    let list_detailtransaksi = []
    let list_detailtransaksi_win = []
    let list_detailtransaksi_notewin = ""
    let list_datasend = []
    let info_card_win = []
    let list_combine = [];
    let round_game_all = -1;
    let counter_transaksi = 0;
    let sum_lose = 0;
    let sum_win = 0;
    let c_before = 0;
    let c_after = 0;
    let flag_all = true
    let flag_bet = true
    let flag_fullbet = true
    let flag_deal = false
    let count_bet = 0;
    let point_result = "";
    let point_style_result = "";
    let list_point_id = "";
    let list_point_style = "";
    let card_result_0_id = "NULL"
    let card_result_1_id = "NULL"
    let card_result_2_id = "NULL"
    let card_result_3_id = "NULL"
    let card_result_4_id = "NULL"
    let card_result_5_id = "NULL"
    let card_result_6_id = "NULL"
    let card_result_0_class = ""
    let card_result_1_class = ""
    let card_result_2_class = ""
    let card_result_3_class = ""
    let card_result_4_class = ""
    let card_result_5_class = ""
    let card_result_6_class = ""
    let card_result_0_img = "./"+path_card+"CARD_FLOP.png"
    let card_result_1_img = "./"+path_card+"CARD_FLOP.png"
    let card_result_2_img = "./"+path_card+"CARD_FLOP.png"
    let card_result_3_img = "./"+path_card+"CARD_FLOP.png"
    let card_result_4_img = "./"+path_card+"CARD_FLOP.png"
    let card_result_5_img = "./"+path_card+"CARD_FLOP.png"
    let card_result_6_img = "./"+path_card+"CARD_FLOP.png"
    let card_result_0_val = ""
    let card_result_1_val = ""
    let card_result_2_val = ""
    let card_result_3_val = ""
    let card_result_4_val = ""
    let card_result_5_val = ""
    let card_result_6_val = ""
    let card_result_array_id = []
    let card_result_array_val = []
    let info_result = "";
    let info_card = [];
    let isModalMinBet = false;
    let isModal_allinvoice = false;
    let isModal_carabermain = false;
    let isModal_detailtransaksi = false;
    let flag_win = false
    let shuffleArray = [];
    let usedIndexes = [];
    $: {
          setInterval(updateClock, 100);
         
    }
    const call_play = () => {
      idtransaksi = "";
      tipe_click = "";
      flag_minimalbet = false;
      c_before = 0;
      c_after = 0;
      count_bet = 0;
      totalbet = 0;
      flag_bet = true;
      flag_fullbet = true;
      flag_deal = false;
      flag_win = false;
      shuffleArray = [];
      usedIndexes = [];
      card_result_array_id = []
      card_result_array_val = []
      info_result = "";
      info_card = [];
      info_card_win = [];
      list_point_style = ""
      list_point_id = ""
      point_result = "";
      point_style_result = "";
      card_result_0_img = "./"+path_card+"CARD_FLOP.png"
      card_result_1_img = "./"+path_card+"CARD_FLOP.png"
      card_result_2_img = "./"+path_card+"CARD_FLOP.png"
      card_result_3_img = "./"+path_card+"CARD_FLOP.png"
      card_result_4_img = "./"+path_card+"CARD_FLOP.png"
      card_result_5_img = "./"+path_card+"CARD_FLOP.png"
      card_result_6_img = "./"+path_card+"CARD_FLOP.png"
      card_result_0_val = ""
      card_result_1_val = ""
      card_result_2_val = ""
      card_result_3_val = ""
      card_result_4_val = ""
      card_result_5_val = ""
      card_result_6_val = ""
      card_result_0_class = ""
      card_result_1_class = ""
      card_result_2_class = ""
      card_result_3_class = ""
      card_result_4_class = ""
      card_result_5_class = ""
      card_result_6_class = ""
      list_detailtransaksi = []
      list_detailtransaksi_win = []
      list_detailtransaksi_notewin = ""
    };
    const call_bet = () => {
      flag_minimalbet = true;
      if(parseInt(min_bet) > 0){
        if(credit > 0){
          if(parseInt(credit) >= parseInt(min_bet)){
            count_bet = count_bet + 1
            totalbet = totalbet + 1
            sound = 0;
            spin[sound].play();
            switch(count_bet){
              case 1:
                flag_deal = true;
                flag_fullbet = false;
                c_before = credit;
                credit = credit - parseInt(min_bet);
                c_after = credit;
                point_style_result = "text-error font-bold"
                point_result = "-" + (totalbet * parseInt(min_bet))
                break;
              case 2:
                c_before = credit;
                credit = credit - parseInt(min_bet);
                c_after = credit;
                point_style_result = "text-error font-bold"
                point_result = "-" + (totalbet * parseInt(min_bet))
                break;
              case 3:
                c_before = credit;
                credit = credit - parseInt(min_bet);
                c_after = credit;
                point_style_result = "text-error font-bold"
                point_result = "-" + (totalbet * parseInt(min_bet))
                break;
              case 4:
                c_before = credit;
                credit = credit - parseInt(min_bet);
                c_after = credit;
                point_style_result = "text-error font-bold"
                point_result = "-" + (totalbet * parseInt(min_bet))
                break;
            }
            factory_click("SINGLE_BET");
           
          }else{
            alert("not enough Credit")
          }
          
        }else{
          alert("Credit Empty")
        }
      }else{
        alert("Please choose minimal bet")
        flag_minimalbet = false;
      }
    };
    const call_fullbet = () => {
      flag_minimalbet = true;
      if(parseInt(min_bet) > 0){
        let credit_after = parseInt(min_bet) * 4;
        if(credit >= credit_after){
          count_bet = count_bet + 4
          totalbet = totalbet + 4
          sound = 0;
          spin[sound].play();
          c_before = credit;
          credit = credit - (parseInt(min_bet) * totalbet);
          c_after = credit;
          point_style_result = "text-error font-bold"
          point_result = "-" + (totalbet * parseInt(min_bet))
          factory_click("FULL_BET")
        }else{
          alert("not enough Credit")
          flag_minimalbet = false;
        }
      }else{
        alert("Please choose minimal bet")
        flag_minimalbet = false;
      }
    };
    const call_deal = () => {
        count_bet = 4
        sound = 0;
        spin[sound].play();
        factory_click("DEAL")
        
    };
   
    function hitung_statuswinlose(data_array){
      let data_result = [];
     
      data_result = royal_flush_factory(data_array);
      if(!data_result[0]){
        data_result = five_kind_factory(data_array);
        if(!data_result[0]){
          data_result = straight_flush_factory(data_array);
          if(!data_result[0]){
            data_result = fourofkind_factory(data_array);
            if(!data_result[0]){
              data_result = fullhouse_factory(data_array);
              if(!data_result[0]){
                data_result = flush_factory(data_array);
                if(!data_result[0]){
                  data_result = straight_factory(data_array);
                  if(!data_result[0]){
                    data_result = threeofkind_factory(data_array);
                    if(!data_result[0]){
                      data_result = twopair_factory(data_array);
                      if(!data_result[0]){
                        data_result = acepair_factory(data_array);
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
     
      
      return [data_result[0],data_result[1],data_result[2]]
    }
    function royal_flush_factory(data_array){
      let flag_func =false
      let data_win = []
   
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].code_card]){
          counts[data_array[i].code_card] += 1
        }else{
          counts[data_array[i].code_card] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 3){
            if(prop == "S"){
              temp.push(prop + ":" + counts[prop])
            }
        }
      }
      if(temp.length > 0){
        let temp_string = temp[0]
        let temp_result = temp_string.split(":");
        let total_temp = temp_result[1];
        let total_jk = 0;
        let total_card = 0;
        if(temp_result[0] == "S"){
          if(parseInt(total_temp) == 5 || parseInt(total_temp) == 6){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "S" || data_array[i].code_card == "JK"){
                switch(data_array[i].val){
                    case "10":
                      total_jk = total_jk + 1;break;
                    case "J":
                      total_jk = total_jk + 1;break;
                    case "K":
                      total_jk = total_jk + 1;break;
                    case "Q":
                      total_jk = total_jk + 1;break;
                    case "AS":
                      total_jk = total_jk + 1;break;
                    case "JK":
                      total_jk = total_jk + 1;break;
                }
              }
            }
            total_card = total_jk
          }
          if(parseInt(total_temp) == 4){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "S" || data_array[i].code_card == "JK"){
                switch(data_array[i].val){
                    case "10":
                      total_jk = total_jk + 1;break;
                    case "J":
                      total_jk = total_jk + 1;break;
                    case "K":
                      total_jk = total_jk + 1;break;
                    case "Q":
                      total_jk = total_jk + 1;break;
                    case "AS":
                      total_jk = total_jk + 1;break;
                    case "JK":
                      total_jk = total_jk + 1;break;
                }
              }
            }
            total_card = total_jk
          }
          if(parseInt(total_temp) == 3){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "S" || data_array[i].code_card == "JK"){
                switch(data_array[i].val){
                    case "10":
                      total_jk = total_jk + 1;break;
                    case "J":
                      total_jk = total_jk + 1;break;
                    case "K":
                      total_jk = total_jk + 1;break;
                    case "Q":
                      total_jk = total_jk + 1;break;
                    case "AS":
                      total_jk = total_jk + 1;break;
                    case "JK":
                      total_jk = total_jk + 1;break;
                }
              }
            }
            total_card = total_jk
          }
        }
       
   
        if(total_card == 5){
            info_result = "Royal Flush"
            info_card = pattern_stright_10
            flag_func = true;
          
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "S" || data_array[i].code_card == "JK"){
                  data_win.push(data_array[i])
              }
            }
          }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,0];
    }
    function five_kind_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val_display]){
          counts[data_array[i].val_display] += 1
        }else{
          counts[data_array[i].val_display] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 3){
              temp.push(prop + ":" + counts[prop])
          }
      }
      if(temp.length > 0){
        let temp_string = temp[0]
        let temp_result = temp_string.split(":");
        let total_temp = temp_result[1];
        let total_jk = 0;
        let total_card = 0;
        if(parseInt(total_temp) == 4){
          for(let i=0;i<data_array.length;i++){
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
            }
          }
          total_card = parseInt(total_temp) + total_jk
          if(total_card == 5){
            info_result = "5 Of A Kind"
            info_card = pattern_stright_10
            flag_func = true;
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0] || data_array[i].val == "JK"){
                  data_win.push(data_array[i])
                }
              }
            }
            // credit_animation(credit,1,totalbet)
          }
        }
        if(parseInt(total_temp) == 3){
          for(let i=0;i<data_array.length;i++){
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
            }
          }
          total_card = parseInt(total_temp) + total_jk
          if(total_card == 5){
            info_result = "5 Of A Kind"
            info_card = pattern_stright_10
            flag_func = true;
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0] || data_array[i].val == "JK"){
                  data_win.push(data_array[i])
                }
              }
            }
  
            // credit_animation(credit,1,totalbet)
            
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,1];
    }
    function straight_flush_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].code_card]){
          counts[data_array[i].code_card] += 1
        }else{
          counts[data_array[i].code_card] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 3){
              temp.push(prop + ":" + counts[prop])
          }
      }
      
      if(temp.length > 0){
        for(let z=0;z<temp.length;z++){
          let objdata_final = []
          let temp_string = temp[z]
          let temp_result = temp_string.split(":");
          if(parseInt(temp_result[1]) == 5){
            for(let i=0;i<data_array.length;i++){
                if(temp_result[0] == data_array[i].code_card){
                  objdata_final.push(data_array[i].val_display)
                }
            }
          
            let flag = []
            flag[0] = checkArray(pattern_stright_1,objdata_final)
            flag[1] = checkArray(pattern_stright_2,objdata_final)
            flag[2] = checkArray(pattern_stright_3,objdata_final)
            flag[3] = checkArray(pattern_stright_4,objdata_final)
            flag[4] = checkArray(pattern_stright_5,objdata_final)
            flag[5] = checkArray(pattern_stright_6,objdata_final)
            flag[6] = checkArray(pattern_stright_7,objdata_final)
            flag[7] = checkArray(pattern_stright_8,objdata_final)
            flag[8] = checkArray(pattern_stright_9,objdata_final)
            flag[9] = checkArray(pattern_stright_10,objdata_final)
            for(let i=0;i<flag.length;i++){
              if(flag[i] == true){
                info_result = "STRAIGHT FLUSH"
                info_card = pattern_stright_10
                flag_func = true;
                
              
                for(let i=0;i<data_array.length;i++){
                  temp_string = temp[0]
                  temp_result = temp_string.split(":");
                  if(temp_result[0] == data_array[i].code_card){
                    data_win.push(data_array[i])
                  }
                }
                break;
              }
            }
          }
          if(parseInt(temp_result[1]) == 4){
            for(let i=0;i<data_array.length;i++){
                if(temp_result[0] == data_array[i].code_card){
                  objdata_final.push(data_array[i].val_display)
                }
            }
           
            let flag = []
            flag[0] = checkArray(pattern_4_1,objdata_final)
            flag[1] = checkArray(pattern_4_2,objdata_final)
            flag[2] = checkArray(pattern_4_3,objdata_final)
            flag[3] = checkArray(pattern_4_4,objdata_final)
            flag[4] = checkArray(pattern_4_5,objdata_final)
            flag[5] = checkArray(pattern_4_6,objdata_final)
            flag[6] = checkArray(pattern_4_7,objdata_final)
            flag[7] = checkArray(pattern_4_8,objdata_final)
            flag[8] = checkArray(pattern_4_9,objdata_final)
            flag[9] = checkArray(pattern_4_10,objdata_final)
            flag[10] = checkArray(pattern_4_11,objdata_final)
            
            for(let i=0;i<flag.length;i++){
              if(flag[i] == true){
                let total_jk = 0;
                let total_card = 0;
                for(let i=0;i<data_array.length;i++){
                  if(data_array[i].code_card == "JK"){
                    total_jk = total_jk + 1
                  }
                }
                total_card = parseInt(temp_result[1]) + total_jk
                
                if(parseInt(total_card) == 5){
                  info_result = "STRAIGHT FLUSH"
                  info_card = pattern_stright_10
                  flag_func = true;
                  
                
                  for(let i=0;i<data_array.length;i++){
                    temp_string = temp[0]
                    temp_result = temp_string.split(":");
                    if(temp_result[0] == data_array[i].code_card || data_array[i].code_card == "JK"){
                      data_win.push(data_array[i])
                    }
                  }
                  break;
                }
              }
            }
          }
          if(parseInt(temp_result[1]) == 3){
            for(let i=0;i<data_array.length;i++){
                if(temp_result[0] == data_array[i].code_card){
                  objdata_final.push(data_array[i].val_display)
                }
            }
          
            let flag = []
            flag[0] = checkArray(pattern_3_1,objdata_final)
            flag[1] = checkArray(pattern_3_2,objdata_final)
            flag[2] = checkArray(pattern_3_3,objdata_final)
            flag[3] = checkArray(pattern_3_4,objdata_final)
            flag[4] = checkArray(pattern_3_5,objdata_final)
            flag[5] = checkArray(pattern_3_6,objdata_final)
            flag[6] = checkArray(pattern_3_7,objdata_final)
            flag[7] = checkArray(pattern_3_8,objdata_final)
            flag[8] = checkArray(pattern_3_9,objdata_final)
            flag[9] = checkArray(pattern_3_10,objdata_final)
            flag[10] = checkArray(pattern_3_11,objdata_final)
            flag[11] = checkArray(pattern_3_12,objdata_final)
           
            for(let i=0;i<flag.length;i++){
              if(flag[i] == true){
                let total_jk = 0;
                let total_card = 0;
                for(let i=0;i<data_array.length;i++){
                  if(data_array[i].code_card == "JK"){
                    total_jk = total_jk + 1
                  }
                }
                total_card = parseInt(temp_result[1]) + total_jk
                
                if(parseInt(total_card) == 5){
                  info_result = "STRAIGHT FLUSH"
                  info_card = pattern_stright_10
                  flag_func = true;
                  
                
                  for(let i=0;i<data_array.length;i++){
                    temp_string = temp[0]
                    temp_result = temp_string.split(":");
                    if(temp_result[0] == data_array[i].code_card || data_array[i].code_card == "JK"){
                      data_win.push(data_array[i])
                    }
                  }
                  break;
                }
              }
            }
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,2];
    }
    function fourofkind_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val]){
          counts[data_array[i].val] += 1
        }else{
          counts[data_array[i].val] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 3){
              temp.push(prop + ":" + counts[prop])
          }
      }
      // console.log(temp)
      let total = 0;
      let total_temp = temp.length
      let temp_string = ""
      let temp_result;
      for(let i=0;i<total_temp;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          total = total + parseInt(temp_result[1])
      }
      // console.log(total)
      if(total == 3){//FOUR OF KIND
        
        let flag_jk = false;
        for(let i=0;i<temp.length;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == temp_result[0]){
              data_win.push(data_array[j])
            }
          }
        }
        for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == "JK"){
              data_win.push(data_array[j])
              flag_jk = true
              break;
            }
        }
        if(flag_jk){
          info_result = "FOUR OF KIND"
          info_card = temp
          flag_func = true
        }
        // credit_animation(credit,3,totalbet)
      }
      if(total == 4){//FOUR OF KIND
        info_result = "FOUR OF KIND"
        info_card = temp
        flag_func = true
        
        for(let i=0;i<temp.length;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == temp_result[0]){
              data_win.push(data_array[j])
            }
          }
        }
  
        // credit_animation(credit,3,totalbet)
      }
      if(total == 6){
        let data_baru = []
        for(let i=0;i<temp.length;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          if(parseInt(temp_result[1]) == 3){
            data_baru.push(temp[i])
            
          }
        }
        for(let i=0;i<data_baru.length-1;i++){
          let temp_string2 = data_baru[i]
          let temp_result2 = temp_string2.split(":");
          
          for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == temp_result2[0]){
              data_win.push(data_array[j])
            }
          }
        }
        for(let j=0;j<data_array.length;j++){
            if(data_array[j].val == "JK"){
              data_win.push(data_array[j])
            }
        }
        if(data_win.length == 4){
          info_result = "FOUR OF KIND"
          info_card = temp
          flag_func = true
          
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,3];
    }
    function fullhouse_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val]){
          counts[data_array[i].val] += 1
        }else{
          counts[data_array[i].val] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 2){
            if(prop != "JK"){
              temp.push(prop + ":" + counts[prop])
            }
          }
      }
      // console.log(temp)
      let total = 0;
      let total_all = 0;
      let total_jk = 0;
      let total_temp = temp.length
      let temp_string = ""
      let temp_result;
      for(let i=0;i<total_temp;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          total = total + parseInt(temp_result[1])
      }
     
      if(total == 4){
        for(let i=0;i<temp.length;i++){
            temp_string = temp[i]
            temp_result = temp_string.split(":");
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].val == temp_result[0]){
                data_win.push(data_array[i])
              }
            }
        }
        total = data_win.length
    
        for(let i=0;i<data_array.length;i++){
          if(data_array[i].val == "JK"){
            total_jk = total_jk + 1
            data_win.push(data_array[i])
          }
        }
        
        total_all = total_all + total + total_jk
        if(total_all == 6){
          let popped = data_win.pop();
          total_all = data_win.length
        }

        if(total_all == 5){
          info_result = "FULL HOUSE"
          info_card = temp
          flag_func = true
        } 
      }
      if(total == 5){//FULL HOUSE
        if(temp.length == 2){
          info_result = "FULL HOUSE"
          info_card = temp
          flag_func = true
          
          for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
          }
          // credit_animation(credit,4,totalbet)
        }
      }
      if(total == 6){
        if(temp.length == 2){
          info_result = "FULL HOUSE"
          info_card = temp
          flag_func = true
          
          for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length-1;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
          }
        }else{
          
          let flag_jk = false;
          for(let i=0;i<temp.length-1;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
          }
          // console.log(data_win.length)
          if(data_win.length < 5){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].val == "JK"){
                data_win.push(data_array[i])
                flag_jk = true;
              }
            }
          }
          if(flag_jk){
            info_result = "FULL HOUSE"
            info_card = temp
            flag_func = true
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,4];
    }
    function flush_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].code_card]){
          counts[data_array[i].code_card] += 1
        }else{
          counts[data_array[i].code_card] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 5){
              temp.push(prop + ":" + counts[prop])
          }
      }
      let total = 0;
      let total_temp = temp.length
      let temp_string = ""
      let temp_result;
      
      for(let i=0;i<total_temp;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");
          
          total = total + parseInt(temp_result[1])
      }
     
      if(total_temp == 1){
        if(total == 5){ //FLUSH
          info_result = "FLUSH"
          info_card = temp
          flag_func = true
  
  
          for(let i=0;i<temp.length;i++){
            temp_string = temp[i]
            temp_result = temp_string.split(":");
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == temp_result[0]){
                data_win.push(data_array[i])
              }
            }
          }
        
         
          // credit_animation(credit,5,totalbet)
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,5];
    }
    function straight_factory(data_array){
      let flag_func = false
      let data_win = []
      let objdata_master = []
      for(let i in data_array){
        objdata_master.push(data_array[i].val_display)
      }
      let flag = []
      
      flag[0] = checkArray(pattern_stright_1,objdata_master)
      flag[1] = checkArray(pattern_stright_2,objdata_master)
      flag[2] = checkArray(pattern_stright_3,objdata_master)
      flag[3] = checkArray(pattern_stright_4,objdata_master)
      flag[4] = checkArray(pattern_stright_5,objdata_master)
      flag[5] = checkArray(pattern_stright_6,objdata_master)
      flag[6] = checkArray(pattern_stright_7,objdata_master)
      flag[7] = checkArray(pattern_stright_8,objdata_master)
      flag[8] = checkArray(pattern_stright_9,objdata_master)
      flag[9] = checkArray(pattern_stright_10,objdata_master)

      flag[10] = checkArray(pattern_4_1,objdata_master)
      flag[11] = checkArray(pattern_4_2,objdata_master)
      flag[12] = checkArray(pattern_4_3,objdata_master)
      flag[13] = checkArray(pattern_4_4,objdata_master)
      flag[14] = checkArray(pattern_4_5,objdata_master)
      flag[15] = checkArray(pattern_4_6,objdata_master)
      flag[16] = checkArray(pattern_4_7,objdata_master)
      flag[17] = checkArray(pattern_4_8,objdata_master)
      flag[18] = checkArray(pattern_4_9,objdata_master)
      flag[19] = checkArray(pattern_4_10,objdata_master)
      flag[20] = checkArray(pattern_4_11,objdata_master)
      // console.log(objdata_master)
      // console.log(flag)
      for(let i=0;i<flag.length;i++){
        if(flag[i] == true){
          info_result = "STRAIGHT"
          info_card = pattern_stright_10
        
          switch(i){
            case 0:
              for(let t=0;t<pattern_stright_1.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_1[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 1:
              for(let t=0;t<pattern_stright_2.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_2[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 2:
              for(let t=0;t<pattern_stright_3.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_3[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 3:
              for(let t=0;t<pattern_stright_4.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_4[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 4:
              for(let t=0;t<pattern_stright_5.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_5[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 5:
              for(let t=0;t<pattern_stright_6.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_6[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 6:
              for(let t=0;t<pattern_stright_7.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_7[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 7:
              for(let t=0;t<pattern_stright_8.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_8[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 8:
              for(let t=0;t<pattern_stright_9.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_9[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 9:
              for(let t=0;t<pattern_stright_10.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_stright_10[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 10:
              for(let t=0;t<pattern_4_1.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_1[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 11:
              for(let t=0;t<pattern_4_2.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_2[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 12:
              for(let t=0;t<pattern_4_3.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_3[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 13:
              for(let t=0;t<pattern_4_4.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_4[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 14:
              for(let t=0;t<pattern_4_5.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_5[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 15:
              for(let t=0;t<pattern_4_6.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_6[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 16:
              for(let t=0;t<pattern_4_7.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_7[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 17:
              for(let t=0;t<pattern_4_8.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_8[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 18:
              for(let t=0;t<pattern_4_9.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_9[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 19:
              for(let t=0;t<pattern_4_10.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_10[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
            case 20:
              for(let t=0;t<pattern_4_11.length;t++){
                let temp_data = data_array.find(card => card.val_display == pattern_4_11[t])
                if(temp_data != undefined){
                  data_win.push(temp_data)
                }
              }
              break;
          }
          let total_card = 0;
          total_card = parseInt(data_win.length)
          if(total_card < 5){
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].code_card == "JK"){
                data_win.push(data_array[i])
              }
            }
          }
        
          total_card = parseInt(data_win.length)
          if(total_card == 5){
            flag_func = true;
          }
          
          break;
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,6];
    }
    function threeofkind_factory(data_array){
      let flag_func = false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val_display]){
          counts[data_array[i].val_display] += 1
        }else{
          counts[data_array[i].val_display] = 1
        }
      }
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 2){
              temp.push(prop + ":" + counts[prop])
          }
      }
      
      if(temp.length > 0){
        let temp_string = temp[0]
        let temp_result = temp_string.split(":");
        let total_temp = temp_result[1];
        let total_jk = 0;
        let total_card = 0;
        
        if(parseInt(total_temp) == 3){
          for(let i=0;i<data_array.length;i++){
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
            }
          }
          total_card = parseInt(total_temp) + total_jk
          if(total_card == 4){
            info_result = "3 Of A Kind"
            info_card = pattern_stright_10
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
            }
            flag_func = true;
          }
          if(total_card == 3){
            info_result = "3 Of A Kind"
            info_card = pattern_stright_10
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
            }
  
            // credit_animation(credit,7,totalbet)
            flag_func = true;
          }
        }
        if(parseInt(total_temp) == 2){
          for(let i=0;i<data_array.length;i++){
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
            }
          }
          total_card = parseInt(total_temp) + total_jk
          if(total_card == 3){
            info_result = "3 Of A Kind"
            info_card = pattern_stright_10
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val_display == temp_result[0]){
                  data_win.push(data_array[i])
                }
                if(data_array[i].val_display == "1"){
                  data_win.push(data_array[i])
                }
              }
            }
            // credit_animation(credit,7,totalbet)
            flag_func = true;
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,7]
    }
    function twopair_factory(data_array){
      let flag_func =false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val]){
          counts[data_array[i].val] += 1
        }else{
          counts[data_array[i].val] = 1
        }
      }
     
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 2){
              temp.push(prop + ":" + counts[prop])
          }
      }
      let total = 0;
      let total_temp = temp.length
      let temp_string = ""
      let temp_result;
      for(let i=0;i<total_temp;i++){
          temp_string = temp[i]
          temp_result = temp_string.split(":");    
          total = total + parseInt(temp_result[1])
      }
    
      let flag_two = false
      if(total_temp == 3){
        if(total == 4 || total == 6){
          for(let x=0;x<temp.length;x++){
            temp_result = temp[x].split(":");
            switch(temp_result[0]){
              case "10":
                flag_two = true;break;
              case "J":
                flag_two = true;break;
              case "Q":
                flag_two = true;break;
              case "K":
                flag_two = true;break;
              case "AS":
                flag_two = true;break;
              case "JK":
                flag_two = true;break;
            }
          }
          
          if(flag_two){//2 PAIR
            info_result = "2 PAIR"
            info_card = temp
            flag_func = true
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
            }
          }
        }
      }
      if(total_temp == 2){
        if(total == 4 || total == 6){
          for(let x=0;x<temp.length;x++){
            temp_result = temp[x].split(":");
            switch(temp_result[0]){
              case "10":
                flag_two = true;break;
              case "J":
                flag_two = true;break;
              case "Q":
                flag_two = true;break;
              case "K":
                flag_two = true;break;
              case "AS":
                flag_two = true;break;
              case "JK":
                flag_two = true;break;
            }
          }
          
          if(flag_two){//2 PAIR
            info_result = "2 PAIR"
            info_card = temp
            flag_func = true
  
            for(let i=0;i<temp.length;i++){
              temp_string = temp[i]
              temp_result = temp_string.split(":");
              for(let i=0;i<data_array.length;i++){
                if(data_array[i].val == temp_result[0]){
                  data_win.push(data_array[i])
                }
              }
            }
          }
        }
      }
      if(flag_func == false){
        data_win = [];
      }
      return [flag_func,data_win,8]
    }
    function acepair_factory(data_array){
      let flag_func = false
      let data_win = []
      let counts = []
      for(let i=0;i<data_array.length;i++){
        if(counts[data_array[i].val_display]){
          counts[data_array[i].val_display] += 1
        }else{
          counts[data_array[i].val_display] = 1
        }
      }
      
      let temp = [];
      for(let prop in counts){
        if (counts[prop] >= 2){
              temp.push(prop + ":" + counts[prop])
          }
      }
      let total_as = 0;
      let total_jk = 0;
      let total_card = 0;
      if(temp.length > 0){
        let temp_string = temp[0]
        let temp_result = temp_string.split(":");
        let total_temp = temp_result[1];
        for(let i=0;i<data_array.length;i++){
            if(data_array[i].val_display == temp_result[0]){
              data_win.push(data_array[i])
            }
        }
        for(let i=0;i<data_win.length;i++){
            if(data_win[i].val == "AS"){
              total_as = total_as + 1;
            }
            if(data_win[i].val == "JK"){
              total_jk = total_jk + 1;
            }
        }
        
        if(total_as == 2){ // 2 AS
            info_result = "ACE PAIR"
            info_card = temp
            // credit_animation(credit,9,totalbet)
            flag_func = true;
        }
        if(total_jk == 2){ // 2 JK
            let popped = data_win.pop();
            total_jk = data_win.length
            total_as = 0
            total_card = 0
            console.log(total_jk)
            for(let i=0;i<data_array.length;i++){
              if(data_array[i].val == "AS"){
                total_as = total_as + 1
                data_win.push(data_array[i])
              }
            }
            total_card = total_as + total_jk
            if(total_card == 2){ // 1AS + 1JK
              info_result = "ACE PAIR"
              info_card = temp
              flag_func = true;
            }
        }
      }else{
        for(let i=0;i<data_array.length;i++){
            if(data_array[i].val == "AS"){
              total_as = total_as + 1;
              data_win.push(data_array[i])
            }
            if(data_array[i].code_card == "JK"){
              total_jk = total_jk + 1;
              data_win.push(data_array[i])
            }
        }
        total_card = total_as + total_jk
    
        if(total_card == 2){ // 1 as + 1 jk
            info_result = "ACE PAIR"
            info_card = temp
            flag_func = true;
        }
      }
  
      if(flag_func == false){
        data_win = [];
      }
  
      return [flag_func,data_win,9]
    }
    
    function checkArray(arr_1,arr_2){
      return arr_1.every((val) => arr_2.includes(val))
    }
    function factory_click(e){
      tipe_click = e
      console.log("TIPE CLICK :" + tipe_click)
      let flag_hitung = true;
      if(e != "DEAL"){
        switch(count_bet){
          case 1:
            flag_hitung = false;
            break;
          case 2:
            // shuffleArray_bet()
            flag_hitung = false;
            break;
          case 3:
            // shuffleArray_bet()
            flag_hitung = false;
            break;
          case 4:
            if(e=="FULL_BET"){
              flag_hitung = false;
              // shuffleArray_fullbet()
              // shuffleArray_card(card_result_data)
              // shuffleArray_fullbet()
            }else{
              sendData(totalbet,min_bet,c_before,c_after,0,"","",shuffleArray,"","LOSE")
              shuffleArray_bet()
              flag_hitung = true;
            }
            
            
            
            break;
        }
      }else{
        shuffleArray_deal()
        flag_hitung = true;
      }
      
      
      if(flag_hitung){
        console.log("TOTAL CARD : "+shuffleArray.length)
        if(shuffleArray.length > 0){
          hitung_rumus(e)
        }
      }else{
        sendData(totalbet,min_bet,c_before,c_after,0,"","",shuffleArray,"","LOSE")
      }
    }
    function hitung_rumus(tipe_data){
        console.log("HITUNG RUMUS")
        
        let status = hitung_statuswinlose(shuffleArray)
        console.log("STATUS HITUNG :" + JSON.stringify(status))
        // console.log("CARD :" + JSON.stringify(shuffleArray))
        if(status[0]){
          flag_win = status[0] 
          sound = 0;
          win[sound].play();
          // console.log(status[1])
          // if(tipe_data != "DEAL"){
          //   // sendData(totalbet,min_bet,c_before,c_after,0,0,"",shuffleArray,"","LOSE")
          // }
          credit_animation_factory(credit,totalbet,status[1],status[2])
        }
    }
    function shuffleArray_card(array){
      let i = 0
      while(i<7){
        let randomNumber = Math.floor(Math.random() * array.length)
        if(!usedIndexes.includes(randomNumber)){
          shuffleArray.push(array[randomNumber]);
          usedIndexes.push(randomNumber);
          i++;
        }
      }  
      
      // shuffleArray = [];
      // shuffleArray.push(array[10]);
      // shuffleArray.push(array[39]);
      // shuffleArray.push(array[17]);
      // shuffleArray.push(array[53]);
      // shuffleArray.push(array[26]);
      // shuffleArray.push(array[11]);
      // shuffleArray.push(array[19]);
      // console.log(shuffleArray)
      
      
     
      // ==== ACE PAIR 1 AS AS===
      // shuffleArray = [];
      // shuffleArray.push(array[12]);
      // shuffleArray.push(array[25]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[18]);
      // shuffleArray.push(array[20]);
      // shuffleArray.push(array[30]);
      // shuffleArray.push(array[28]);
      // console.log(shuffleArray)
  
  
      // ==== ACE PAIR 1 AS JK===
      // shuffleArray = [];
      // shuffleArray.push(array[12]);
      // shuffleArray.push(array[53]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[18]);
      // shuffleArray.push(array[20]);
      // shuffleArray.push(array[30]);
      // shuffleArray.push(array[28]);
      // console.log(shuffleArray)
      
  
      // ==== 2 PAIR ===
      // shuffleArray = [];
      // shuffleArray.push(array[8]);
      // shuffleArray.push(array[21]);
      // shuffleArray.push(array[22]);
      // shuffleArray.push(array[48]);
      // shuffleArray.push(array[49]);
      // shuffleArray.push(array[30]);
      // shuffleArray.push(array[28]);
      // console.log(shuffleArray)
  
      // ==== 2 PAIR 1 ===
      // shuffleArray = [];
      // shuffleArray.push(array[25]);
      // shuffleArray.push(array[48]);
      // shuffleArray.push(array[13]);
      // shuffleArray.push(array[29]);
      // shuffleArray.push(array[24]);
      // shuffleArray.push(array[51]);
      // shuffleArray.push(array[16]);
      // console.log(shuffleArray)
  
     
  
      // ==== 3 kind 1 ===
      // shuffleArray = [];
      // shuffleArray.push(array[10]);
      // shuffleArray.push(array[23]);
      // shuffleArray.push(array[36]);
      // shuffleArray.push(array[39]);
      // shuffleArray.push(array[3]);
      // shuffleArray.push(array[18]);
      // shuffleArray.push(array[32]);
      // console.log(shuffleArray)
  
      // ==== 3 kind 2 ===
      // shuffleArray = [];
      // shuffleArray.push(array[10]);
      // shuffleArray.push(array[23]);
      // shuffleArray.push(array[39]);
      // shuffleArray.push(array[3]);
      // shuffleArray.push(array[18]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[53]);
      // console.log(shuffleArray)
  
      // ==== 3 kind 3 ===
      // shuffleArray = [];
      // shuffleArray.push(array[12]);
      // shuffleArray.push(array[18]);
      // shuffleArray.push(array[25]);
      // shuffleArray.push(array[36]);
      // shuffleArray.push(array[52]);
      // shuffleArray.push(array[0]);
      // shuffleArray.push(array[49]);
      // console.log(shuffleArray)
  
      // ==== STRAIGHT ===
      // shuffleArray = [];
      // shuffleArray.push(array[0]);
      // shuffleArray.push(array[14]);
      // shuffleArray.push(array[28]);
      // shuffleArray.push(array[42]);
      // shuffleArray.push(array[4]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[50]);
      // console.log(shuffleArray)
  
      // ==== STRAIGHT 2 ===
      // shuffleArray = [];
      // shuffleArray.push(array[0]);
      // shuffleArray.push(array[14]);
      // shuffleArray.push(array[28]);
      // shuffleArray.push(array[42]);
      // shuffleArray.push(array[4]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[16]);
      // console.log(shuffleArray)
  
      // ==== STRAIGHT 3 ===
      // shuffleArray = [];
      // shuffleArray.push(array[47]);
      // shuffleArray.push(array[35]);
      // shuffleArray.push(array[23]);
      // shuffleArray.push(array[11]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[38]);
      // shuffleArray.push(array[25]);
      // console.log(shuffleArray)
  
  
      // ==== FLUSH ===
      // shuffleArray = [];
      // shuffleArray.push(array[0]);
      // shuffleArray.push(array[2]);
      // shuffleArray.push(array[29]);
      // shuffleArray.push(array[6]);
      // shuffleArray.push(array[8]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[4]);
      // console.log(shuffleArray)
  
      // ==== FULL HOUSE ===
      // shuffleArray = [];
      // shuffleArray.push(array[29]);
      // shuffleArray.push(array[14]);
      // shuffleArray.push(array[27]);
      // shuffleArray.push(array[6]);
      // shuffleArray.push(array[8]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[1]);
      // console.log(shuffleArray)
  
      // ==== FULL HOUSE 2===
      // shuffleArray = [];
      // shuffleArray.push(array[35]);
      // shuffleArray.push(array[3]);
      // shuffleArray.push(array[23]);
      // shuffleArray.push(array[9]);
      // shuffleArray.push(array[21]);
      // shuffleArray.push(array[48]);
      // shuffleArray.push(array[42]);
      // console.log(shuffleArray)
  
      // ==== FULL HOUSE ===
      // shuffleArray = [];
      // shuffleArray.push(array[1]);
      // shuffleArray.push(array[14]);
      // shuffleArray.push(array[27]);
      // shuffleArray.push(array[6]);
      // shuffleArray.push(array[8]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[29]);
      // console.log(shuffleArray)
  
      // ==== 4 OF A KIND ===
      // shuffleArray = [];
      // shuffleArray.push(array[6]);
      // shuffleArray.push(array[32]);
      // shuffleArray.push(array[45]);
      // shuffleArray.push(array[8]);
      // shuffleArray.push(array[31]);
      // shuffleArray.push(array[19]);
      // shuffleArray.push(array[29]);
      // console.log(shuffleArray)
  
      
  
      // ==== STRAIGHT FLUSH ===
      // shuffleArray = [];
      // shuffleArray.push(array[0]);
      // shuffleArray.push(array[1]);
      // shuffleArray.push(array[2]);
      // shuffleArray.push(array[3]);
      // shuffleArray.push(array[4]);
      // shuffleArray.push(array[31]);
      // shuffleArray.push(array[29]);
      // console.log(shuffleArray)
  
      
      // ==== 5 OF A KIND 1 ===
      // shuffleArray = [];
      // shuffleArray.push(array[0]);
      // shuffleArray.push(array[13]);
      // shuffleArray.push(array[26]);
      // shuffleArray.push(array[39]);
      // shuffleArray.push(array[52]);
      // shuffleArray.push(array[31]);
      // shuffleArray.push(array[29]);
      // console.log(shuffleArray)
  
      
      // ==== 5 OF A KIND 2 ===
      // shuffleArray = [];
      // shuffleArray.push(array[0]);
      // shuffleArray.push(array[13]);
      // shuffleArray.push(array[26]);
      // shuffleArray.push(array[53]);
      // shuffleArray.push(array[52]);
      // shuffleArray.push(array[31]);
      // shuffleArray.push(array[29]);
      // console.log(shuffleArray)
  
    
      // ==== ROYAL FLUSH ===
      // shuffleArray = [];
      // shuffleArray.push(array[47]);
      // shuffleArray.push(array[48]);
      // shuffleArray.push(array[49]);
      // shuffleArray.push(array[50]);
      // shuffleArray.push(array[51]);
      // shuffleArray.push(array[31]);
      // shuffleArray.push(array[29]);
      // console.log(shuffleArray)
  
      console.log(shuffleArray)
    }
    function shuffleArray_bet(){
      if(count_bet == 4){
        flag_bet = false
      }
  
      if(count_bet == 1){
        card_result_0_id = shuffleArray[0].id
        card_result_2_id = shuffleArray[2].id
        card_result_0_img = shuffleArray[0].img
        card_result_1_img = "./"+path_card+"CARD_FLOP.png"
        card_result_2_img = shuffleArray[2].img
        card_result_0_val = shuffleArray[0].val
        card_result_1_val = "NULL"
        card_result_2_val = shuffleArray[2].val
      }
      if(count_bet == 2){
        card_result_4_id = shuffleArray[4].id
        card_result_3_img = "./"+path_card+"CARD_FLOP.png"
        card_result_4_img = shuffleArray[4].img
        card_result_3_val = "NULL"
        card_result_4_val = shuffleArray[4].val
       
      }
      if(count_bet == 3){
        card_result_5_id = shuffleArray[5].id
        card_result_5_img = shuffleArray[5].img
        card_result_5_val = shuffleArray[5].val
      }
      if(count_bet == 4){
        card_result_1_id = shuffleArray[1].id
        card_result_3_id = shuffleArray[3].id
        card_result_6_id = shuffleArray[6].id
        card_result_1_img = shuffleArray[1].img
        card_result_3_img = shuffleArray[3].img
        card_result_6_img = shuffleArray[6].img
        card_result_1_val = shuffleArray[1].val
        card_result_3_val = shuffleArray[3].val
        card_result_6_val = shuffleArray[6].val
  
        card_result_array_id.push(card_result_0_id)
        card_result_array_id.push(card_result_1_id)
        card_result_array_id.push(card_result_2_id)
        card_result_array_id.push(card_result_3_id)
        card_result_array_id.push(card_result_4_id)
        card_result_array_id.push(card_result_5_id)
        card_result_array_id.push(card_result_6_id)
  
        card_result_array_val.push(card_result_0_val)
        card_result_array_val.push(card_result_1_val)
        card_result_array_val.push(card_result_2_val)
        card_result_array_val.push(card_result_3_val)
        card_result_array_val.push(card_result_4_val)
        card_result_array_val.push(card_result_5_val)
        card_result_array_val.push(card_result_6_val)
  
        // hitung(card_result_array_id,card_result_array_val);
      }
      return shuffleArray;
    }
    function shuffleArray_fullbet(){
      if(count_bet == 4){
        flag_bet = false
      }
      card_result_0_id = shuffleArray[0].id
      card_result_1_id = shuffleArray[1].id
      card_result_2_id = shuffleArray[2].id
      card_result_3_id = shuffleArray[3].id
      card_result_4_id = shuffleArray[4].id
      card_result_5_id = shuffleArray[5].id
      card_result_6_id = shuffleArray[6].id
      card_result_0_img = shuffleArray[0].img
      card_result_1_img = shuffleArray[1].img
      card_result_2_img = shuffleArray[2].img
      card_result_3_img = shuffleArray[3].img
      card_result_4_img = shuffleArray[4].img
      card_result_5_img = shuffleArray[5].img
      card_result_6_img = shuffleArray[6].img
      card_result_0_val = shuffleArray[0].val
      card_result_1_val = shuffleArray[1].val
      card_result_2_val = shuffleArray[2].val
      card_result_3_val = shuffleArray[3].val
      card_result_4_val = shuffleArray[4].val
      card_result_5_val = shuffleArray[5].val
      card_result_6_val = shuffleArray[6].val
  
      card_result_array_id.push(card_result_0_id)
      card_result_array_id.push(card_result_1_id)
      card_result_array_id.push(card_result_2_id)
      card_result_array_id.push(card_result_3_id)
      card_result_array_id.push(card_result_4_id)
      card_result_array_id.push(card_result_5_id)
      card_result_array_id.push(card_result_6_id)
  
      card_result_array_val.push(card_result_0_val)
      card_result_array_val.push(card_result_1_val)
      card_result_array_val.push(card_result_2_val)
      card_result_array_val.push(card_result_3_val)
      card_result_array_val.push(card_result_4_val)
      card_result_array_val.push(card_result_5_val)
      card_result_array_val.push(card_result_6_val)

     
    }
    function shuffleArray_deal(){
      if(count_bet == 4){
        flag_bet = false
      }
     
      card_result_0_id = shuffleArray[0].id
      card_result_1_id = shuffleArray[1].id
      card_result_2_id = shuffleArray[2].id
      card_result_3_id = shuffleArray[3].id
      card_result_4_id = shuffleArray[4].id
      card_result_5_id = shuffleArray[5].id
      card_result_6_id = shuffleArray[6].id
      card_result_0_img = shuffleArray[0].img
      card_result_1_img = shuffleArray[1].img
      card_result_2_img = shuffleArray[2].img
      card_result_3_img = shuffleArray[3].img
      card_result_4_img = shuffleArray[4].img
      card_result_5_img = shuffleArray[5].img
      card_result_6_img = shuffleArray[6].img
      card_result_0_val = shuffleArray[0].val
      card_result_1_val = shuffleArray[1].val
      card_result_2_val = shuffleArray[2].val
      card_result_3_val = shuffleArray[3].val
      card_result_4_val = shuffleArray[4].val
      card_result_5_val = shuffleArray[5].val
      card_result_6_val = shuffleArray[6].val
  
      card_result_array_id.push(card_result_0_id)
      card_result_array_id.push(card_result_1_id)
      card_result_array_id.push(card_result_2_id)
      card_result_array_id.push(card_result_3_id)
      card_result_array_id.push(card_result_4_id)
      card_result_array_id.push(card_result_5_id)
      card_result_array_id.push(card_result_6_id)
  
      card_result_array_val.push(card_result_0_val)
      card_result_array_val.push(card_result_1_val)
      card_result_array_val.push(card_result_2_val)
      card_result_array_val.push(card_result_3_val)
      card_result_array_val.push(card_result_4_val)
      card_result_array_val.push(card_result_5_val)
      card_result_array_val.push(card_result_6_val)
      
      
      // c_before = credit;
      // credit = credit - (parseInt(min_bet) * totalbet);
      // c_after = credit;
    }
    
  
    function credit_animation_factory(credit_before,total_bet,data_win,n){
      let point = (list_point_db[n].lispoin_poin* total_bet)*parseInt(min_bet);
      point_result = "+" + point
      let credit_target = credit_before + ((list_point_db[n].lispoin_poin* total_bet)*parseInt(min_bet))
      
      animateValue(credit, credit_target, 500)
      point_style_result = "text-secondary font-bold";
      list_point_style = "bg-slate-500"
      list_point_id = list_point_db[n].lispoin_id
      console.log("POINT ID :" + list_point_id)
      c_before = credit;
      credit = credit - (parseInt(min_bet) * totalbet);
      c_after = credit;
  
      card_result_0_class = "brightness-50"
      card_result_1_class = "brightness-50"
      card_result_2_class = "brightness-50"
      card_result_3_class = "brightness-50"
      card_result_4_class = "brightness-50"
      card_result_5_class = "brightness-50"
      card_result_6_class = "brightness-50"
      
      card_result_0_id = shuffleArray[0].id
      card_result_1_id = shuffleArray[1].id
      card_result_2_id = shuffleArray[2].id
      card_result_3_id = shuffleArray[3].id
      card_result_4_id = shuffleArray[4].id
      card_result_5_id = shuffleArray[5].id
      card_result_6_id = shuffleArray[6].id

     
     
    for(let j=0;j<data_win.length;j++){
          let flag_data = true;
          var result = shuffleArray.find(item => item.id === data_win[j].id);
          if(result.id != ""){
            flag_data = false
          }
          
          // console.log("card result 0 " + result.id)
          // console.log("card result 1 " + card_result_0_id)
          switch(result.id){
            case card_result_0_id:
              if(!flag_data){
                card_result_0_class = ""
              }
              break;
            case card_result_1_id:
              if(!flag_data){
                card_result_1_class = ""
              }
              break;
            case card_result_2_id:
              if(!flag_data){
                card_result_2_class = ""
              }
              break;
            case card_result_3_id:
              if(!flag_data){
                card_result_3_class = ""
              }
              break;
            case card_result_4_id:
              if(!flag_data){
                card_result_4_class = ""
              }
              break;
            case card_result_5_id:
              if(!flag_data){
                card_result_5_class = ""
              }
              break;
            case card_result_6_id:
              if(!flag_data){
                card_result_6_class = ""
              }
              break;
          }
     }
        
       
      
      sendData(total_bet,0,c_before,credit_target,point,list_point_db[n].lispoin_id,info_result,shuffleArray,data_win,"WIN")
  
      flag_all = false
    }
   
    function animateValue(start, end, duration) {
      // assumes integer values for start and end
      
     
      var range = end - start;
      // no timer shorter than 50ms (not really visible any way)
      var minTimer = 50;
      // calc step time to show all interediate values
      var stepTime = Math.abs(Math.floor(duration / range));
      
      // never go below minTimer
      stepTime = Math.max(stepTime, minTimer);
      
      // get current time and calculate desired end time
      var startTime = new Date().getTime();
      var endTime = startTime + duration;
      var timer;
    
      function run() {
          var now = new Date().getTime();
          var remaining = Math.max((endTime - now) / duration, 0);
          var value = Math.round(end - (remaining * range));
          credit = value;
          if (value == end) {
              clearInterval(timer);
              flag_all = true;
          }
      }
      
      timer = setInterval(run, stepTime);
      run();
    }
    function sendData(data_roundbet,data_minbet,data_cbefore,data_cafter,data_win,code_win,note_win,data_resultcard,data_resultcardwin,data_statustransaksi){
      let status_css = "";
      if(data_statustransaksi == "LOSE"){
        status_css = "bg-error text-white";
        sum_lose = sum_lose + (parseInt(data_roundbet) * parseInt(data_minbet));
      }else{
        status_css = "bg-success text-black";
        sum_win = sum_win + parseInt(data_win);
      }
      counter_transaksi = counter_transaksi + 1
      const objSend = {
        id_transaksi: dayjs().tz(client_timezone).format("YYYYMMDDHHmmss")+counter_transaksi,
        date_transaksi:clockmachine_data,
        round_bet:parseInt(data_roundbet),
        bet: parseInt(data_minbet), 
        total_bet: parseInt(data_roundbet) * parseInt(data_minbet), 
        credit_before: parseInt(data_cbefore), 
        credit_after: parseInt(data_cafter), 
        win: data_win, 
        code_win: code_win, 
        note_win: note_win, 
        result_card: data_resultcard,
        result_card_win: data_resultcardwin,
        status_transaction_css: status_css,
        status_transaction: data_statustransaksi
      };
      list_datasend = [...list_datasend,objSend]
      
      list_datasend.sort((a, b) => b.id_transaksi - a.id_transaksi);

      let resultwin = "";
      let total_datawin = data_resultcardwin.length
      for(let i=0; i<total_datawin; i++){
        if(i==total_datawin-1){
          resultwin += data_resultcardwin[i].id
        }else{
          resultwin += data_resultcardwin[i].id + ","
        }
      }
     
      if(data_roundbet == 1 || data_roundbet == 4){ 
        if(idtransaksi == ""){
          savetransaksi(data_roundbet,data_minbet,data_cbefore,data_cafter,data_win,code_win,resultwin,data_statustransaksi)
        }else{
          savetransaksidetail(idtransaksi,data_roundbet,data_minbet,data_cbefore,data_cafter,data_win,code_win,resultwin,data_statustransaksi)
        }
      }else{
        if(idtransaksi == ""){
          savetransaksi(data_roundbet,data_minbet,data_cbefore,data_cafter,data_win,code_win,resultwin,data_statustransaksi)
        }else{
          savetransaksidetail(idtransaksi,data_roundbet,data_minbet,data_cbefore,data_cafter,data_win,code_win,resultwin,data_statustransaksi)
        }
        
      }

    }
    const handleInformation = () => {
        if(!flag_minimalbet){
          isModalMinBet = true
        }
    };
    const handle_minbet = (e) => {
        list_point_db = []
        for(let i=0;i<client_listbet.length;i++){
          if(client_listbet[i].lisbet_minbet == parseInt(e)){
            min_bet = client_listbet[i].lisbet_minbet
            for(let j=0;j<client_listbet[i].lisbet_config.length;j++){
              list_point_db = [
                  ...list_point_db,
                  {
                    lispoin_id: client_listbet[i].lisbet_config[j].lispoin_id,
                    lispoin_nmpoin: client_listbet[i].lisbet_config[j].lispoin_nmpoin,
                    lispoin_poin: client_listbet[i].lisbet_config[j].lispoin_poin,
                  },
              ];
            }
          }
        }

        isModalMinBet = false
    };
    const call_allinvoice = () => {
      isModal_allinvoice = true
      fetch_invoicell()
    };
    const call_detailinvoice = (e,f,d) => {
      isModal_detailtransaksi = true
      // alert(f)
      list_detailtransaksi = []
      list_detailtransaksi_win = []
      list_detailtransaksi_notewin = d

     
      const card_result = e.split("-");
      const card_win = f.split(",");
      for(let i = 0; i < card_result.length; i++) {
        list_detailtransaksi.push(card_result_data[card_result[i]].img);
      }
      for(let i = 0; i < card_win.length; i++) {
        for(let j = 0; j < card_result_data.length; j++) {
          if(card_result_data[j].id === card_win[i]){
            list_detailtransaksi_win.push(card_result_data[j].img)
          }
        }
      }

    };
    const call_carabermain = () => {
      isModal_carabermain = true
    };
    async function savetransaksi(c_roundbet,c_bet,c_before,c_after,c_win,c_idpoin,resultcardwin,c_status) {
      flag_all = false
      round_game_all = round_game_all + 1
      const res = await fetch(path_api+"api/savetransaksi", {
          method: "POST",
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify({
            transaksi_company: client_company,
            transaksi_username: client_username,
            transaksi_roundgameall: parseInt(round_game_all),
            transaksi_roundbet: parseInt(c_roundbet),
            transaksi_bet: parseInt(c_bet),
            transaksi_cbefore: parseInt(c_before),
            transaksi_cafter: parseInt(c_after),
            transaksi_win: parseInt(c_win),
            transaksi_codepoin: c_idpoin,
            transaksi_resultcardwin: resultcardwin,
            transaksi_status: c_status,
          }),
      });
      const json = await res.json();
      if (json.status === 400) {
          // logout();
      } else if (json.status == 403) {
          alert(json.message);
      } else {
      
        idtransaksi = json.client_idtransaksi
      
        let card = "8-34-23-47-42-40-15"
        // let card = json.client_cardgame
        let card_length = parseInt(json.client_cardlength)
        const myArray = card.split("-");
        shuffleArray = [];
        for(let i = 0; i < myArray.length; i++) {
          shuffleArray.push(card_result_data[myArray[i]]);
        }
        if((card_length-1) == parseInt(round_game_all)){
          round_game_all = -1
        }
        if(c_roundbet == 4){
          // console.log("TOTAL CARD : "+shuffleArray.length)
          hitung_rumus(tipe_click)
          if(tipe_click == "FULL_BET") {
            shuffleArray_fullbet()
          }else{
            shuffleArray_bet()
          }
        }else{
          shuffleArray_bet()
        }

        
        flag_all = true;
      }
    }
    async function savetransaksidetail(c_idtransaksi,c_roundbet,c_bet,c_before,c_after,c_win,c_idpoin,resultcardwin,c_status) {
      flag_all = false
      // console.log(c_idpoin)
      const res = await fetch(path_api+"api/savetransaksidetail", {
          method: "POST",
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify({
            transaksidetail_company: client_company,
            transaksidetail_idtransaksi: c_idtransaksi,
            transaksidetail_roundbet: parseInt(c_roundbet),
            transaksidetail_bet: parseInt(c_bet),
            transaksidetail_cbefore: parseInt(c_before),
            transaksidetail_cafter: parseInt(c_after),
            transaksidetail_win: parseInt(c_win),
            transaksidetail_codepoin: c_idpoin,
            transaksidetail_resultcardwin: resultcardwin,
            transaksidetail_status: c_status,
          }),
      });
      const json = await res.json();
      if (json.status === 400) {
          // logout();
      } else if (json.status == 403) {
          alert(json.message);
      } else {
        if(tipe_click == "FULL_BET") {
          shuffleArray_fullbet()
        }else{
          shuffleArray_bet()
        }
        flag_all = true;
      }
    }
    async function fetch_invoicell() {
      list_invoice = []
      const res = await fetch(path_api+"api/listinvoice", {
          method: "POST",
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify({
            invoice_company: client_company,
            Invoice_username: client_username,
          }),
      });
      const json = await res.json();
      if (json.status === 400) {
          // logout();
      } else if (json.status == 403) {
          alert(json.message);
      } else {
        let record = json.record;
        if (record != null) {
            for (var i = 0; i < record.length; i++) {
              let winlose = parseInt(record[i]["invoice_totalwin"]) - (parseInt(record[i]["invoice_round"]*record[i]["invoice_totalbet"]))
              let winlose_css = "";
              let status_css = ""
              if(winlose >0){
                winlose_css = "text-accent";
              }else{
                winlose_css = "text-error";
              }
              if(record[i]["invoice_status"] == "LOSE"){
                status_css = "bg-error text-white";
              }else{
                status_css = "bg-success text-black";
              }
              list_invoice = [
                    ...list_invoice,
                    {
                        invoice_id: record[i]["invoice_id"],
                        invoice_date: record[i]["invoice_date"],
                        invoice_round: record[i]["invoice_round"],
                        invoice_totalbet: record[i]["invoice_totalbet"],
                        invoice_totalwin: record[i]["invoice_totalwin"],
                        invoice_winlose: winlose,
                        invoice_winlose_css: winlose_css,
                        invoice_card_result: record[i]["invoice_card_result"],
                        invoice_card_win: record[i]["invoice_card_win"],
                        invoice_nmpoin: record[i]["invoice_nmpoin"],
                        invoice_status: record[i]["invoice_status"],
                        invoice_status_css: status_css,
                    },
                ];
            }
        }
      }
    }
  </script>
  
  
    <div class="navbar">
      <div class="navbar-start">
        <a href="/?token={token_browser}" >
            <img src="https://sdsb4d.com/logo-green.svg" alt="SDSB" class="hover:scale-110  transition ">
        </a>
      </div>
      <div class="navbar-center hidden lg:flex">
        <div class="flex flex-row gap-2">
            <label on:click={() => {
              call_allinvoice();
            }} class="btn bg-base-300 border-none shadow-lg shadow-green-500/50 btn-sm">INVOICE</label>
            <label on:click={() => {
              call_carabermain();
            }} class="btn bg-base-300 border-none shadow-lg shadow-green-500/50 btn-sm">CARA BERMAIN</label>
        </div>
      </div>
      <div class="navbar-end hidden text-xs lg:text-sm lg:inline-block text-right">
          <div class="flex items-start  ">
            <div class="flex flex-col w-full">
              <p class="w-full text-xs lg:text-sm text-right select-none">
                Asia/Jakarta <br />
                {clockmachine}  WIB (+7)<br>
                {client_name} <br />
                {client_ipaddress}
              </p>
              <div class="w-full text-xs lg:text-sm text-right select-none">
                CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(credit)}</span>
                <span class="{point_style_result}">{point_result}</span>
              </div>
              <div class="w-full text-xs lg:text-sm text-right select-none">
                ROUND BET {totalbet}x<span class="link-accent">{min_bet}</span> : <span class="text-error">{new Intl.NumberFormat().format(min_bet*totalbet)}</span>
              </div>
            </div>
          </div>
      </div>
    </div>
  
    <section class="flex lg:hidden justify-center w-full mb-2 gap-2 bg-base-200 p-2 rounded-md ">
      <p class="w-full text-xs  text-left">
        Asia/Jakarta <br />
        {clockmachine}  WIB (+7)<br>
        developer <br />
        {client_ipaddress}
      </p>
      <p class="w-full text-xs  text-right">
        CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(credit)}</span>
        <span class="{point_style_result}">{point_result}</span>
        <br />
        ROUND BET {totalbet}x : <span class="text-error">{new Intl.NumberFormat().format(min_bet*totalbet)}</span>
      </p>
    </section>
  
    <section class="flex lg:hidden justify-center w-full mb-3 gap-2">
      <div on:click={() => {
        call_allinvoice();
      }} class="btn bg-base-300 border-none shadow-lg shadow-green-500/50 btn-xs">INVOICE</div>
      <div on:click={() => {
        call_carabermain();
      }} class="btn bg-base-300 border-none shadow-lg shadow-green-500/50 btn-xs">CARA BERMAIN</div>
    </section>
  
    <section class="w-full select-none rounded-md p-2 mt-2 bg-base-100  ">
      <div class="grid grid-cols-2 w-full ">
   
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[0]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[0]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[0]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[5]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[5]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[5]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[1]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[1]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[1]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[6]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[6]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[6]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[2]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[2]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[2]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[7]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[7]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[7]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[3]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[3]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[3]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[8]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[8]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[8]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full pr-2 lg:pr-5 {list_point_db[4]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[4]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[4]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
        <div class="flex w-full  pr-2 lg:pr-5 {list_point_db[9]["lispoin_id"] == list_point_id ? list_point_style:''}">
          <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[9]["name"]}</div>
          <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point_db[9]["lispoin_poin"]*totalbet)*parseInt(min_bet))}</div>
        </div>
      </div>
      
    </section>
    
    
    <section class="w-full flex justify-center mt-5 ">
      <div class="avatar-group -space-x-14 lg:-space=x=6">
        <div class="relative">
          <img class="{card_result_0_class}" src="{card_result_0_img}" >
        </div>
        <div class="relative">
          <img class="{card_result_1_class}" src="{card_result_1_img}" >
        </div>
        <div class="relative">
          <img class="{card_result_2_class}" src="{card_result_2_img}" >
        </div>
        <div class="relative">
          <img class="{card_result_3_class}" src="{card_result_3_img}" >
        </div>
        <div class="relative">
          <img class="{card_result_4_class}" src="{card_result_4_img}" >
        </div>
        <div class="relative">
          <img class="{card_result_5_class}" src="{card_result_5_img}" >
        </div>
        <div class="relative">
          <img class="{card_result_6_class}" src="{card_result_6_img}" >
        </div>
      </div>
    </section>
   
    <section class="flex flex-col lg:flex-row w-full mt-5 gap-2">
      <div class="flex w-full p-0 justify-center lg:justify-end">
        <div class="flex flex-col w-1/3 ">
          <span class="text-sm lg:text-lg w-full text-center select-none">Minimal Bet</span>
          <div class=" bg-base-300 text-lg p-2 w-full cursor-pointer text-center rounded-lg" on:click={() => {
              handleInformation();
            }}>
            <span class="text-center link-accent select-none">{new Intl.NumberFormat().format(min_bet)}</span>
          </div>
        </div>
      </div>
      <div class="flex mt-1 lg:mt-0  lg:place-self-end gap-1 w-full ">
        {#if flag_all}
          {#if !flag_bet}
            <button on:click={() => {
                call_play();
              }} class="btn btn-success btn-md w-full lg:w-[240px] ">Play</button>
          {/if}
          {#if flag_bet}
            <button on:click={() => {
                call_bet();
              }} class="btn btn-primary btn-md w-1/2 lg:w-[120px] ">BET</button>
            {#if flag_fullbet}
            <button on:click={() => {
                call_fullbet();
              }} class="btn btn-primary btn-md w-1/2 lg:w-[120px]" >FULL BET</button>
            {/if}
            {#if flag_deal}
            <button on:click={() => {
                call_deal();
              }} class="btn btn-primary btn-md w-1/2 lg:w-[120px]">DEAL</button>
            {/if}
          {/if}
        {/if}
      </div>
    </section>
   
  
  

  
  
  
  <input type="checkbox" id="my-modal-information" class="modal-toggle" bind:checked={isModalMinBet}>
  <div class="modal" on:click|self={()=>isModalMinBet = false}>
      <div class="modal-box relative w-11/12 max-w-lg h-1/2 lg:h-1/4 overflow-hidden select-none">
          <label for="my-modal-information" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
          <h3 class="text-xs lg:text-sm font-bold -mt-2">MINIMAL BET</h3>
          <div class="h-fit overflow-auto scrollbar-hide mt-2" >
              <div class="grid grid-cols-5 mt-5 gap-2 justify-self-center">
                {#each client_listbet as rec}
                  <div
                    on:click={() => {
                      handle_minbet(rec.lisbet_minbet);
                    }} 
                    class="btn btn-sm btn-outline btn-success cursor-pointer">{new Intl.NumberFormat().format(rec.lisbet_minbet)}</div>
                {/each}
              </div>
              
          </div>
      </div>
  </div>
  
  <input type="checkbox" id="my-modal-allinvoice" class="modal-toggle" bind:checked={isModal_allinvoice}>
  <div class="modal" on:click|self={()=>isModal_allinvoice = false}>
    <div class="modal-box relative select-none  lg:max-w-4xl h-full lg:max-h-[600px] rounded-none lg:rounded-lg p-2 lg:p-4 overflow-hidden">
      <label for="my-modal-allinvoice" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
      <h3 class="text-xs lg:text-sm font-bold -mt-1">INVOICE</h3>
      <div class="overflow-auto h-[90%] scrollbar-thin scrollbar-thumb-green-100 mt-4">
          <table class="table table-xs w-full" >
              <thead class="sticky top-0">
                  <tr>
                    <th colspan=3 class="text-center">Update Every 3 Minute</th>
                  </tr>
                  <tr>
                      <th width="5%" class="text-xs text-center align-top">STATUS</th>
                      <th width="5%" class="text-xs text-left align-top">INVOICE TRANSAKSI</th>
                      <th width="*" class="text-xs text-right align-top">WINLOSE</th>
                  </tr>
              </thead>
              <tbody>
                  {#each list_invoice as rec}
                  <tr>
                    <td class="text-xs  text-center whitespace-nowrap align-top">
                      <span class="{rec.invoice_status_css} p-1 text-xs lg:text-sm  uppercase  rounded-lg w-20 ">{rec.invoice_status}</span>
                    </td>
                    <td on:click={() => {
                      call_detailinvoice(rec.invoice_card_result,rec.invoice_card_win,rec.invoice_nmpoin);
                      }} class="text-xs  text-left whitespace-nowrap align-top cursor-pointer underline">
                      {rec.invoice_id}<br />
                      {rec.invoice_date}<br />
                    </td>
                    <td class="text-xs text-right  whitespace-nowrap align-top {rec.invoice_winlose_css}">{new Intl.NumberFormat().format(rec.invoice_winlose)}</td>
                  </tr>
                  {/each}
              </tbody>
          </table>
      </div>
    </div>
  </div>
  
  <input type="checkbox" id="my-modal-carabermain" class="modal-toggle" bind:checked={isModal_carabermain}>
  <div class="modal" on:click|self={()=>isModal_carabermain = false}>
    <div class="modal-box relative select-none  lg:max-w-2xl h-full lg:max-h-[600px] rounded-none lg:rounded-lg p-2 lg:p-4 overflow-hidden">
      <label for="my-modal-carabermain" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
      <h3 class="text-xs lg:text-sm font-bold -mt-1">CARA BERMAIN</h3>
      <div class="overflow-auto h-[90%] scrollbar-thin scrollbar-thumb-green-100 mt-4">
          <p class="text-sm">
            Mickey Mouse Bolatangkas adalah permainan dengan menggunakan tujuh buah kartu yang akan dibagikan secara random pada setiap putaran permainan 
            sehingga pemain bisa mendapatkan susunan kartu yang ditentukan dan tertinggi untuk mencapai kemenangan. sebelum bermain silahkan login dan masuk ke 
            permainan Mickey Mouse Bolatangkas terlebih dahulu.<br />
            Ini adalah tampilan tab Mickey Mouse Bolatangkas. Berikut beberapa tombol/panel yang wajib anda ketahui<br />
  
            Setting, pilih tombol ini jika anda ingin mengubah pengaturan Game<br />
            Info Panel, disini anda bisa melihat info user id, jumlah koin anda, nomor meja.<br />
            Coin in, klik untuk menambahkan koin ke credit anda<br />
            Coin out, klik untuk memindahkan koin dari credit ke dompet anda<br />
            Fullhouse, klik untuk mengubah pilihan fullhouse anda<br />
            Info bet<br />
            Reward info<br />
            Klik untuk memilih room permainan, setiap room memiliki jumlah taruhan yang berbeda<br />
            Open, klik untuk membuka kartu<br />
            Fold, klik untung membuang kartu<br /><br />
            Sesuai dengan nilai kartu yang didapatkan apabila pemain menang berikut adalah susunan set kartu dari yang terkecil hingga yang tertinggi
            <br /><br />
            - Royal flush : kombinasi dari kartu ace, king, queen, jack, dan 10 dengan jenis yang sama<br />
            - 5 of a kind : kombinasi dari four of a kind dan kartu joker<br />
            - Straight flush : kombinasi dari straight dan flush<br />
            - 4 of a kind : kombinasi dari 4 kartu yang sama<br />
            - Full house : kombinasi dari 3 kartu dan 2 kartu yang sama<br />
            - Flush : kombinasi dari 5 kartu yang memiliki jenis yang sama<br />
            - Straight : kombinasi 5 kartu yang berurutan<br />
            - 3 of a kind : kombinasi dari 3 kartu yang sama<br />
            - 2 pair ( 10 pair ) : kombinasi dari 2 pasang kartu yang sama dengan catatan nilai kartu diatas 10<br />
            - Ace pair : kombinasi dari 2 kartu ace yang sama<br />
          </p>
      </div>
    </div>
  </div>
  
  <input type="checkbox" id="my-modal-detailtransaksi" class="modal-toggle" bind:checked={isModal_detailtransaksi}>
  <div class="modal" on:click|self={()=>isModal_detailtransaksi = false}>
    <div class="modal-box relative select-none  lg:max-w-4xl h-full lg:max-h-[600px] rounded-none lg:rounded-lg p-2 lg:p-4 overflow-hidden">
      <label for="my-modal-detailtransaksi" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
      <h3 class="text-xs lg:text-sm font-bold -mt-1">DETAIL</h3>
      <div class="overflow-auto h-[90%] scrollbar-thin scrollbar-thumb-green-100 mt-4">
        <div class="flex flex-col gap-1 w-full">
          <div class="grid grid-cols-7 w-full mt-5">
            {#each list_detailtransaksi as rec2}
              <img  src="{rec2}"> 
            {/each}
          </div>
          <div class="w-full text-center">
            {list_detailtransaksi_notewin}
          </div>
          <div class="grid grid-cols-7 mt-5 w-full">
            {#each list_detailtransaksi_win as rec2}
            <img  src="{rec2}" > 
            {/each}
          </div>
        </div>
      </div>
    </div>
  </div>
  

  
  