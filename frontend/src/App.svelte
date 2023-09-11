<script>
  import dayjs from "dayjs";
  import utc from "dayjs/plugin/utc";
  import timezone from "dayjs/plugin/timezone";
  
  dayjs.extend(utc);
  dayjs.extend(timezone);
  let client_ipaddress = "123.456.234.123";
  let client_timezone = "Asia/Jakarta";
  let clockmachine = "";
  let clockmachine_data = "";
  let credit = 50000
  let list_min_bet = [100,500,1000,1500,5000,10000,25000]
  let min_bet = "100"
  let totalbet = 0
  let sound = 0;
  let flag_minimalbet = false;
  let path_card = "CARD/WHITE/";
  let spin = [
    new Audio("/sounds/spin.mp3"),
  ];
  let win = [
    new Audio("/sounds/win.mp3"),
    new Audio("/sounds/win.mp3"),
  ];
  let list_point = [
    {id:"1",code:"RF",name:"Royal Flush",poin:500},
    {id:"2",code:"5K",name:"5 Of A Kind",poin:200},
    {id:"3",code:"SF",name:"Straight Flush",poin:120},
    {id:"4",code:"4K",name:"4 Of A Kind",poin:50},
    {id:"5",code:"FH",name:"Full House",poin:7},
    {id:"6",code:"FL",name:"Flush",poin:5},
    {id:"7",code:"ST",name:"Straight",poin:3},
    {id:"8",code:"3K",name:"3 Of A Kind",poin:2},
    {id:"9",code:"2P",name:"2 Pair (10 PAIR)",poin:1},
    {id:"10",code:"AP",name:"Ace Pair",poin:1},
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
  
  let list_datasend = []
  let info_card_win = []
  let counter_transaksi = 0;
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
  let flag_win = false
  let shuffleArray = [];
  let usedIndexes = [];
  $: {
        setInterval(updateClock, 100);
       
  }
  const call_play = () => {
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
              shuffleArray_card(card_result_data)
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
          
          shuffleArray_bet()
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
        shuffleArray_card(card_result_data)
        shuffleArray_fullbet()
      }else{
        alert("Credit Empty")
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
      shuffleArray_deal()
	};
  function hitung(arr_id,arr_val){
    let flag_royalflush = false
    let flag_fiveofkind = false
    let flag_straight_flush = false
    let flag_fourofkind = false
    let flag_fullhouse = false
    let flag_flush = false
    let flag_straight = false
    let flag_threeofkind = false
    let flag_twopair = false
    
    flag_royalflush = royal_flush(arr_id)
    if(!flag_royalflush){
      flag_fiveofkind = five_kind(arr_id)
      if(!flag_fiveofkind){
        flag_straight_flush = straight_flush(arr_id)
        if(!flag_straight_flush){
          flag_fourofkind = fourofkind(arr_id,arr_val)
          if(!flag_fourofkind){
            flag_fullhouse = fullhouse(arr_id,arr_val)
            if(!flag_fullhouse){
              flag_flush = flush(arr_id,arr_val)
              if(!flag_flush){
                flag_straight = straight(arr_id)
                if(!flag_straight){
                  flag_threeofkind = threeofkind(arr_id)
                  if(!flag_threeofkind){
                    flag_twopair = twopair(arr_id,arr_val)
                    if(!flag_twopair){
                      acepair(arr_id)
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
    if(flag_win){
      sound = 0;
      win[sound].play();
    }
  }
  function royal_flush(arr_id){
    let flag_func = false
    let obj = []
    let objdata_tampung_win = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["val"] = temp_data.val 
      obj["code_card"] = temp_data.code_card 
      obj["val_display"] = temp_data.val_display 
      obj["img"] = temp_data.img 
      objdata_master.push(obj)
      obj = []
    }
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].code_card]){
        counts[objdata_master[i].code_card] += 1
      }else{
        counts[objdata_master[i].code_card] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 4){
            temp.push(prop + ":" + counts[prop])
        }
    }
    // console.log(temp)
    if(temp.length > 0){
      let temp_string = temp[0]
      let temp_result = temp_string.split(":");
      let total_temp = temp_result[1];
      let total_jk = 0;
      let total_card = 0;
      if(parseInt(total_temp) == 5){
        for(let i=0;i<objdata_master.length;i++){
          switch(objdata_master[i].val){
            case "10":
              obj["id"] = objdata_master[i].id
              obj["img"] = objdata_master[i].img
              objdata_tampung_win.push(obj)
              obj = [];
              total_jk = total_jk + 1;break;
            case "J":
              obj["id"] = objdata_master[i].id
              obj["img"] = objdata_master[i].img
              objdata_tampung_win.push(obj)
              obj = [];
              total_jk = total_jk + 1;break;
            case "K":
              obj["id"] = objdata_master[i].id
              obj["img"] = objdata_master[i].img
              objdata_tampung_win.push(obj)
              obj = [];
              total_jk = total_jk + 1;break;
            case "Q":
              obj["id"] = objdata_master[i].id
              obj["img"] = objdata_master[i].img
              objdata_tampung_win.push(obj)
              obj = [];
              total_jk = total_jk + 1;break;
            case "AS":
              obj["id"] = objdata_master[i].id
              obj["img"] = objdata_master[i].img
              objdata_tampung_win.push(obj)
              obj = [];
              total_jk = total_jk + 1;break;
            case "JK":
              obj["id"] = objdata_master[i].id
              obj["img"] = objdata_master[i].img
              objdata_tampung_win.push(obj)
              obj = [];
              total_jk = total_jk + 1;break;
          }
        }
        total_card = total_jk
        if(total_card == 5){
          info_result = "Royal Flush"
          info_card = pattern_stright_10
          flag_win = true
          
          
          for(let x=0;x<objdata_tampung_win.length;x++){
            obj["id"] = objdata_tampung_win[x].id
            obj["img"] = objdata_tampung_win[x].img
            info_card_win.push(obj)
            obj = [];
          }

          credit_animation(credit,0,totalbet)
          flag_func = true;
        }
      }
    }
    return flag_func
  }
  function five_kind(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["code_card"] = temp_data.code_card 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(obj)
      obj = []
    }
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].val_display]){
        counts[objdata_master[i].val_display] += 1
      }else{
        counts[objdata_master[i].val_display] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 3){
            temp.push(prop + ":" + counts[prop])
        }
    }
    // console.log(temp)
    if(temp.length > 0){
      let temp_string = temp[0]
      let temp_result = temp_string.split(":");
      let total_temp = temp_result[1];
      let total_jk = 0;
      let total_card = 0;
      if(parseInt(total_temp) == 4){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        total_card = parseInt(total_temp) + total_jk
        if(total_card == 5){
          info_result = "5 Of A Kind"
          info_card = pattern_stright_10
          flag_win = true


          for(let i in arr_id){
            let temp_data = card_result_data.find(card => card.id == arr_id[i])
            // console.log(temp_data.id+"-"+temp_data.img+"-"+temp_data.val_display)
            for(let x=0;x<temp.length;x++){
              temp_result = temp[x].split(":");
              if(temp_result[0] == temp_data.val_display){
                obj["id"] = temp_data.id
                obj["img"] = temp_data.img
                info_card_win.push(obj)
                obj = [];
              }
              if("1" == temp_data.val_display){
                obj["id"] = temp_data.id
                obj["img"] = temp_data.img
                info_card_win.push(obj)
                obj = [];
              }
            }
          }

          credit_animation(credit,1,totalbet)
          flag_func = true;
        }
      }
      if(parseInt(total_temp) == 3){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        total_card = parseInt(total_temp) + total_jk
        if(total_card == 5){
          info_result = "5 Of A Kind"
          info_card = pattern_stright_10
          flag_win = true

          for(let i in arr_id){
            let temp_data = card_result_data.find(card => card.id == arr_id[i])
            // console.log(temp_data.id+"-"+temp_data.img+"-"+temp_data.val_display)
            for(let x=0;x<temp.length;x++){
              temp_result = temp[x].split(":");
              if(temp_result[0] == temp_data.val_display){
                obj["id"] = temp_data.id
                obj["img"] = temp_data.img
                info_card_win.push(obj)
                obj = [];
              }
              if("1" == temp_data.val_display){
                obj["id"] = temp_data.id
                obj["img"] = temp_data.img
                info_card_win.push(obj)
                obj = [];
              }
            }
          }

          credit_animation(credit,1,totalbet)
          flag_func = true;
        }
      }
    }
    return flag_func
  }
  function straight_flush(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["code_card"] = temp_data.code_card 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(obj)
      obj = []
    }
    
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].code_card]){
        counts[objdata_master[i].code_card] += 1
      }else{
        counts[objdata_master[i].code_card] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 5){
            temp.push(prop + ":" + counts[prop])
        }
    }
    if(temp.length > 0){
      let obj_final = []
      let objdata_final = []
      let temp_string = ""
      let temp_result;
      for(let i=0;i<arr_id.length;i++){
          let temp_data = card_result_data.find(card => card.id == arr_id[i])
          temp_string = temp[0]
          temp_result = temp_string.split(":");
          if(temp_result[0] == temp_data.code_card){
            obj_final["id"] = temp_data.id 
            obj_final["code_card"] = temp_data.code_card 
            obj_final["val_display"] = temp_data.val_display 
            objdata_final.push(temp_data.val_display)
            obj_final = []
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
          flag_win = true

          let temp_data_win = []
          for(let k in arr_id){
              let temp_data = card_result_data.find(card => card.id == arr_id[k])            
              switch(i){
                case 0:
                  for(let t=0;t<pattern_stright_1.length;t++){
                    if(pattern_stright_1[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 1:
                  for(let t=0;t<pattern_stright_2.length;t++){
                    if(pattern_stright_2[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 2:
                  for(let t=0;t<pattern_stright_3.length;t++){
                    if(pattern_stright_3[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 3:
                  for(let t=0;t<pattern_stright_4.length;t++){
                    if(pattern_stright_4[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 4:
                  for(let t=0;t<pattern_stright_5.length;t++){
                    if(pattern_stright_5[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 5:
                  for(let t=0;t<pattern_stright_6.length;t++){
                    if(pattern_stright_6[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 6:
                  for(let t=0;t<pattern_stright_7.length;t++){
                    if(pattern_stright_7[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 7:
                  for(let t=0;t<pattern_stright_8.length;t++){
                    if(pattern_stright_8[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 8:
                  for(let t=0;t<pattern_stright_9.length;t++){
                    if(pattern_stright_9[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
                case 9:
                  for(let t=0;t<pattern_stright_10.length;t++){
                    if(pattern_stright_10[t] == temp_data.val_display){
                      obj["id"] = temp_data.id
                      obj["img"] = temp_data.img
                      obj["val_display"] = temp_data.val_display
                      temp_data_win.push(obj)
                      obj = [];
                    }
                  }
                  break;
              }
          }

          const filteredArr = temp_data_win.reduce((temp_data_win, current) => {
            const x = temp_data_win.find(item => item.val_display === current.val_display);
            if (!x) {
              return temp_data_win.concat([current]);
            } else {
              return temp_data_win;
            }
          }, []);

      
          for(let t=0;t<filteredArr.length;t++){
              obj["id"] = filteredArr[t].id
              obj["img"] = filteredArr[t].img
              info_card_win.push(obj)
              obj = [];
          }


          credit_animation(credit,2,totalbet)
          flag_func = true;
          break;
        }
      }
    }
    return flag_func;
  }
  function fourofkind(arr_id,arr_val){
    let flag =false
    let counts = []
    let obj = []
    let data_result_id = []
    let data_result = []
    for(let i=0;i<arr_val.length;i++){
      if(counts[arr_val[i]]){
        counts[arr_val[i]] += 1
      }else{
        counts[arr_val[i]] = 1
      }
    }
    let temp = [];
    for(let prop in counts){
      if (counts[prop] >= 4){
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
        obj["val"] = temp_result[0]
        obj["count"] = temp_result[1]
        for(let x=0;x<arr_val.length;x++){
          if(temp_result[0] == arr_val[x]){
            data_result_id.push(arr_id[x])
          }
        }
        obj["id"] = data_result_id
        data_result.push(obj)
        obj = []
        data_result_id = []
        total = total + parseInt(temp_result[1])
    }
    if(total == 4){//FOUR OF KIND
      info_result = "FOUR OF KIND"
      info_card = temp
      flag_win = true
      flag = true
      

      for(let i in arr_id){
        let temp_data = card_result_data.find(card => card.id == arr_id[i])
        // console.log(temp_data.id+"-"+temp_data.img+"-"+temp_data.val_display)
        for(let x=0;x<temp.length;x++){
          temp_result = temp[x].split(":");
          if(temp_result[0] == temp_data.val_display){
            obj["id"] = temp_data.id
            obj["img"] = temp_data.img
            info_card_win.push(obj)
            obj = [];
          }
        }
      }

      credit_animation(credit,3,totalbet)
    }
    return flag
  }
  function fullhouse(arr_id,arr_val){
    let flag =false
    let counts = []
    let obj = []
    let data_result_id = []
    let data_result = []
    for(let i=0;i<arr_val.length;i++){
      if(counts[arr_val[i]]){
        counts[arr_val[i]] += 1
      }else{
        counts[arr_val[i]] = 1
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
        obj["val"] = temp_result[0]
        obj["count"] = temp_result[1]
        for(let x=0;x<arr_val.length;x++){
          if(temp_result[0] == arr_val[x]){
            data_result_id.push(arr_id[x])
          }
        }
        obj["id"] = data_result_id
        data_result.push(obj)
        obj = []
        data_result_id = []
        total = total + parseInt(temp_result[1])
    }
    if(total == 5){//FULL HOUSE
      if(temp.length == 2){
        info_result = "FULL HOUSE"
        info_card = temp
        flag_win = true
        flag = true
      
        for(let i in arr_id){
          let temp_data = card_result_data.find(card => card.id == arr_id[i])
          // console.log(temp_data.id+"-"+temp_data.img+"-"+temp_data.val_display)
          for(let x=0;x<temp.length;x++){
            temp_result = temp[x].split(":");
            if(temp_result[0] == temp_data.val){
              obj["id"] = temp_data.id
              obj["img"] = temp_data.img
              info_card_win.push(obj)
              obj = [];
            }
          }
        }
        credit_animation(credit,4,totalbet)
      }
    }
    return flag
  }
  function flush(arr_id,arr_val){
    let flag =false
    let counts = []
    let obj = []
    let data_result_id = []
    let data_result = []
    for(let i=0;i<arr_val.length;i++){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      if(counts[temp_data.code_card]){
        counts[temp_data.code_card] += 1
      }else{
        counts[temp_data.code_card] = 1
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
        obj["val"] = temp_result[0]
        obj["count"] = temp_result[1]
        for(let x=0;x<arr_val.length;x++){
          if(temp_result[0] == arr_val[x]){
            data_result_id.push(arr_id[x])
          }
        }
        obj["id"] = data_result_id
        data_result.push(obj)
        obj = []
        data_result_id = []
        total = total + parseInt(temp_result[1])
    }
    if(total_temp == 1){
      if(total == 5){ //FLUSH
        info_result = "FLUSH"
        info_card = temp
        flag_win = true
        flag = false

        
        for(let i in arr_id){
          let temp_data = card_result_data.find(card => card.id == arr_id[i])
          // console.log(temp_data.id+"-"+temp_data.img+"-"+temp_data.code_card)
          for(let x=0;x<temp.length;x++){
            temp_result = temp[x].split(":");
            if(temp_result[0] == temp_data.code_card){
              obj["id"] = temp_data.id
              obj["img"] = temp_data.img
              info_card_win.push(obj)
              obj = [];
            }
          }
        }

        credit_animation(credit,5,totalbet)
      }
    }
    return flag
  }
  function straight(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(temp_data.val_display)
      obj = []
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
    for(let i=0;i<flag.length;i++){
      if(flag[i] == true){
        info_result = "STRAIGHT"
        info_card = pattern_stright_10
        flag_win = true

        let temp_data_win = []
        for(let k in arr_id){
            let temp_data = card_result_data.find(card => card.id == arr_id[k])            
            switch(i){
              case 0:
                for(let t=0;t<pattern_stright_1.length;t++){
                  if(pattern_stright_1[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 1:
                for(let t=0;t<pattern_stright_2.length;t++){
                  if(pattern_stright_2[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 2:
                for(let t=0;t<pattern_stright_3.length;t++){
                  if(pattern_stright_3[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 3:
                for(let t=0;t<pattern_stright_4.length;t++){
                  if(pattern_stright_4[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 4:
                for(let t=0;t<pattern_stright_5.length;t++){
                  if(pattern_stright_5[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 5:
                for(let t=0;t<pattern_stright_6.length;t++){
                  if(pattern_stright_6[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 6:
                for(let t=0;t<pattern_stright_7.length;t++){
                  if(pattern_stright_7[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 7:
                for(let t=0;t<pattern_stright_8.length;t++){
                  if(pattern_stright_8[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 8:
                for(let t=0;t<pattern_stright_9.length;t++){
                  if(pattern_stright_9[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
              case 9:
                for(let t=0;t<pattern_stright_10.length;t++){
                  if(pattern_stright_10[t] == temp_data.val_display){
                    obj["id"] = temp_data.id
                    obj["img"] = temp_data.img
                    obj["val_display"] = temp_data.val_display
                    temp_data_win.push(obj)
                    obj = [];
                  }
                }
                break;
            }
        }

        const filteredArr = temp_data_win.reduce((temp_data_win, current) => {
          const x = temp_data_win.find(item => item.val_display === current.val_display);
          if (!x) {
            return temp_data_win.concat([current]);
          } else {
            return temp_data_win;
          }
        }, []);

     
        for(let t=0;t<filteredArr.length;t++){
            obj["id"] = filteredArr[t].id
            obj["img"] = filteredArr[t].img
            info_card_win.push(obj)
            obj = [];
        }
        
        credit_animation(credit,6,totalbet)
        flag_func = true;
        break;
      }
    }
    return flag_func;
  }
  function threeofkind(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["code_card"] = temp_data.code_card 
      obj["val_display"] = temp_data.val_display 
      objdata_master.push(obj)
      obj = []
    }
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].val_display]){
        counts[objdata_master[i].val_display] += 1
      }else{
        counts[objdata_master[i].val_display] = 1
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
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        total_card = parseInt(total_temp) + total_jk
        if(total_card == 3){
          info_result = "3 Of A Kind"
          info_card = pattern_stright_10
          flag_win = true

          for(let i in arr_id){
            let temp_data = card_result_data.find(card => card.id == arr_id[i])
            // console.log(temp_data.id+"-"+temp_data.img+"-"+temp_data.val_display)
            for(let x=0;x<temp.length;x++){
              temp_result = temp[x].split(":");
              if(temp_result[0] == temp_data.val_display){
                obj["id"] = temp_data.id
                obj["img"] = temp_data.img
                info_card_win.push(obj)
                obj = [];
              }
            }
          }

          credit_animation(credit,7,totalbet)
          flag_func = true;
        }
      }
      if(parseInt(total_temp) == 2){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
          }
        }
        total_card = parseInt(total_temp) + total_jk
    
        if(total_card == 3){
          info_result = "3 Of A Kind"
          info_card = pattern_stright_10
          flag_win = true

          for(let i in arr_id){
            let temp_data = card_result_data.find(card => card.id == arr_id[i])
            // console.log(temp_data.id+"-"+temp_data.img+"-"+temp_data.val_display)
            for(let x=0;x<temp.length;x++){
              temp_result = temp[x].split(":");
              if(temp_result[0] == temp_data.val_display || "1" == temp_data.val_display){
                obj["id"] = temp_data.id
                obj["img"] = temp_data.img
                info_card_win.push(obj)
                obj = [];
              }
            }
          }

          credit_animation(credit,7,totalbet)
          flag_func = true;
        }
      }
    }
    return flag_func
  }
  function twopair(arr_id,arr_val){
    let flag =false
    let counts = []
    let obj = []
    let data_result_id = []
    let data_result = []
    for(let i=0;i<arr_val.length;i++){
      if(counts[arr_val[i]]){
        counts[arr_val[i]] += 1
      }else{
        counts[arr_val[i]] = 1
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
        obj["val"] = temp_result[0]
        obj["count"] = temp_result[1]
        for(let x=0;x<arr_val.length;x++){
          if(temp_result[0] == arr_val[x]){
            data_result_id.push(arr_id[x])
          }
        }
        obj["id"] = data_result_id
        data_result.push(obj)
        obj = []
        data_result_id = []
        total = total + parseInt(temp_result[1])
    }
    let flag_two = false
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
          flag_win = true
          flag = true

          for(let i in arr_id){
            let temp_data = card_result_data.find(card => card.id == arr_id[i])
            for(let x=0;x<info_card.length;x++){
              temp_result = info_card[x].split(":");
              if(temp_result[0] == temp_data.val){
                obj["id"] = temp_data.id
                obj["img"] = temp_data.img
                info_card_win.push(obj)
                obj = [];
              }
            }
          }
          credit_animation(credit,8,totalbet)
        }
      }
    }
    return flag
  }
  function acepair(arr_id){
    let flag_func = false
    let obj = []
    let objdata_master = []
    for(let i in arr_id){
      let temp_data = card_result_data.find(card => card.id == arr_id[i])
      obj["id"] = temp_data.id 
      obj["code_card"] = temp_data.code_card 
      obj["val"] = temp_data.val 
      obj["val_display"] = temp_data.val_display 
      obj["img"] = temp_data.img 
      objdata_master.push(obj)
      obj = []
    }
    let counts = []
    for(let i=0;i<objdata_master.length;i++){
      if(counts[objdata_master[i].val_display]){
        counts[objdata_master[i].val_display] += 1
      }else{
        counts[objdata_master[i].val_display] = 1
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
      
      if(parseInt(total_temp) == 2){
        for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].val == "AS"){
            total_as = total_as + 1;
            obj["id"] = objdata_master[i].id
            obj["img"] = objdata_master[i].img
            info_card_win.push(obj)
            obj = []
          }
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
            obj["id"] = objdata_master[i].id
            obj["img"] = objdata_master[i].img
            info_card_win.push(obj)
            obj = []
          }
        }
        
        
        if(total_as == 2){ // 2 AS
          info_result = "ACE PAIR"
          info_card = temp
          flag_win = true
          credit_animation(credit,9,totalbet)
          flag_func = true;
        }
      }
    }else{
      for(let i=0;i<objdata_master.length;i++){
          if(objdata_master[i].val == "AS"){
            total_as = total_as + 1;
            obj["id"] = objdata_master[i].id
            obj["img"] = objdata_master[i].img
            info_card_win.push(obj)
            obj = []
          }
          if(objdata_master[i].code_card == "JK"){
            total_jk = total_jk + 1;
            obj["id"] = objdata_master[i].id
            obj["img"] = objdata_master[i].img
            info_card_win.push(obj)
            obj = []
          }
      }
      total_card = total_as + total_jk
      if(total_card == 2){ // 1 as + 1 jk
          info_result = "ACE PAIR"
          info_card = temp
          flag_win = true
          credit_animation(credit,9,totalbet)
          flag_func = true;
      }
    }
    return flag_func
  }
  function checkArray(arr_1,arr_2){
    return arr_1.every((val) => arr_2.includes(val))
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
    // ==== ACE PAIR ===
    // shuffleArray = [];
    // shuffleArray.push(array[12]);
    // shuffleArray.push(array[25]);
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

    // ==== FLUSH ===
    // shuffleArray = [];
    // shuffleArray.push(array[0]);
    // shuffleArray.push(array[2]);
    // shuffleArray.push(array[4]);
    // shuffleArray.push(array[6]);
    // shuffleArray.push(array[8]);
    // shuffleArray.push(array[32]);
    // shuffleArray.push(array[29]);
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
    // shuffleArray.push(array[19]);
    // shuffleArray.push(array[32]);
    // shuffleArray.push(array[45]);
    // shuffleArray.push(array[8]);
    // shuffleArray.push(array[31]);
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
  }
  function shuffleArray_bet(){
    if(count_bet == 4){
      flag_bet = false
    }
    sendData(totalbet,min_bet,c_before,c_after,0,0,"",shuffleArray,"","LOSE")
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
      // console.log(card_result_array_id)
      hitung(card_result_array_id,card_result_array_val);
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

   

  
    sendData(totalbet,min_bet,c_before,c_after,0,0,"",shuffleArray,"","LOSE")
    hitung(card_result_array_id,card_result_array_val);
    return shuffleArray;
  }
  function shuffleArray_deal(){
    if(count_bet == 4){
      flag_bet = false
    }
    // console.log("total bet : " + totalbet)
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
    
    
    c_before = credit;
    credit = credit - (parseInt(min_bet) * totalbet);
    c_after = credit;
    hitung(card_result_array_id,card_result_array_val);
  }
  function credit_animation(credit_before,n,total_bet){
    let point = (list_point[n].poin* total_bet)*parseInt(min_bet);
    // console.log("POINT : " + point)
    point_result = "+" + point
    let credit_target = credit_before + ((list_point[n].poin* total_bet)*parseInt(min_bet))
    
    animateValue(credit, credit_target, 500)
    point_style_result = "text-secondary font-bold";
    list_point_style = "bg-slate-500"
    list_point_id = list_point[n].id

    
    // console.log("POINT ID : " + list_point_id)
    // console.log("INFO RESULT : " + info_result)
    // console.log("INFO CARD : " + info_card)


    c_before = credit;
    credit = credit - (parseInt(min_bet) * totalbet);
    c_after = credit;

    // console.log(shuffleArray)
    // console.log(info_card_win)
    // console.log(card_result_0_id+","+card_result_1_id+","+card_result_2_id+","+card_result_3_id+","+card_result_4_id+","+card_result_5_id+","+card_result_6_id)
    // for(let i in info_card_win){
    //   // let temp_data = shuffleArray.find(card => card.id != info_card_win[i].id)
    //   switch(info_card_win[i].id){
    //     case card_result_0_id:
    //       card_result_0_class = "brightness-50"
    //       break;
    //     case card_result_1_id:
    //       card_result_1_class = "brightness-50"
    //       break;
    //     case card_result_2_id:
    //       card_result_2_class = "brightness-50"
    //       break;
    //     case card_result_3_id:
    //       card_result_3_class = "brightness-50"
    //       break;
    //     case card_result_4_id:
    //       card_result_4_class = "brightness-50"
    //       break;
    //     case card_result_5_id:
    //       card_result_5_class = "brightness-50"
    //       break;
    //     case card_result_6_id:
    //       card_result_6_class = "brightness-50"
    //       break;
    //   }
    // }
    card_result_0_class = "brightness-50"
    card_result_1_class = "brightness-50"
    card_result_2_class = "brightness-50"
    card_result_3_class = "brightness-50"
    card_result_4_class = "brightness-50"
    card_result_5_class = "brightness-50"
    card_result_6_class = "brightness-50"

    for(let i=0;i<shuffleArray.length;i++){
      let flag_data = true;
      for(let j=0;j<info_card_win.length;j++){
        if(shuffleArray[i].id == info_card_win[j].id){
          flag_data = false
        }
      }
      switch(shuffleArray[i].id){
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
    
    sendData(0,0,c_before,credit_target,point,0,info_result,shuffleArray,info_card_win,"WIN")

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
    }else{
      status_css = "bg-success text-black";
    }
    counter_transaksi = counter_transaksi + 1
    const objSend = {
      id_transaksi:counter_transaksi,
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
  }
  const handleInformation = () => {
      if(!flag_minimalbet){
        isModalMinBet = true
      }
  };
  const handle_minbet = (e) => {
      min_bet = parseInt(e)
      isModalMinBet = false
  };
  const call_allinvoice = () => {
    isModal_allinvoice = true
    // fetch_invoicell()
	};
  const call_carabermain = () => {
    isModal_carabermain = true
	};
</script>

<main class="container mx-auto px-2 text-base-content glass xl:rounded-box xl:mt-7 max-w-screen-xl bg-opacity-60 pb-5 h-fit lg:h-full">
  <div class="navbar">
    <div class="navbar-start">
      <a href="/?token=" >
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
            <p class="w-full text-xs lg:text-sm text-right">
              Asia/Jakarta <br />
              {clockmachine}  WIB (+7)<br>
              developer <br />
              {client_ipaddress}
            </p>
            <div class="w-full text-xs lg:text-sm text-right">
              CREDIT : IDR <span class="link-accent" style="--value:15;">{new Intl.NumberFormat().format(credit)}</span>
              <span class="{point_style_result}">{point_result}</span>
            </div>
            <div class="w-full text-xs lg:text-sm text-right ">
              ROUND BET {totalbet}x<span class="link-accent">{min_bet}</span> : <span class="text-error">{new Intl.NumberFormat().format(min_bet*totalbet)}</span>
            </div>
          </div>
        </div>
    </div>
  </div>

  <section class="flex lg:hidden justify-center w-full mb-2 gap-2 bg-base-200 p-2 rounded-md">
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
      <div class="flex w-full pr-2 lg:pr-5 {list_point[0]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[0]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[0]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full pr-2 lg:pr-5 {list_point[5]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[5]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[5]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full pr-2 lg:pr-5 {list_point[1]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[1]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[1]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full pr-2 lg:pr-5 {list_point[6]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[6]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[6]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full pr-2 lg:pr-5 {list_point[2]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[2]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[2]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full pr-2 lg:pr-5 {list_point[7]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[7]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[7]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full pr-2 lg:pr-5 {list_point[3]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[3]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[3]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full pr-2 lg:pr-5 {list_point[8]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[8]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[8]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full pr-2 lg:pr-5 {list_point[4]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[4]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[4]["poin"]*totalbet)*parseInt(min_bet))}</div>
      </div>
      <div class="flex w-full  pr-2 lg:pr-5 {list_point[9]["id"] == list_point_id ? list_point_style:''}">
        <div class="w-full text-xs lg:text-sm whitespace-nowrap">{list_point[9]["name"]}</div>
        <div class="w-full text-xs lg:text-sm whitespace-nowrap text-right link-accent">{new Intl.NumberFormat().format((list_point[9]["poin"]*totalbet)*parseInt(min_bet))}</div>
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
 
</main>

<footer class="footer footer-center p-4 text-base-content mt-2 text-center">
  <div class="grid">
    <p class="text-xs lg:text-sm text-center">
      0.0.1
      <br />
      PowerBy
    </p>
    <img src="https://sdsb4d.com/logo-green.svg" alt="SDSB" class="w-24 lg:w-28">
  </div>
</footer>

<input type="checkbox" id="my-modal-information" class="modal-toggle" bind:checked={isModalMinBet}>
<div class="modal" on:click|self={()=>isModalMinBet = false}>
    <div class="modal-box relative w-11/12 max-w-lg h-1/2 lg:h-1/4 overflow-hidden select-none">
        <label for="my-modal-information" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
        <h3 class="text-xs lg:text-sm font-bold -mt-2">MINIMAL BET</h3>
        <div class="h-fit overflow-auto scrollbar-hide mt-2" >
            <div class="grid grid-cols-5 mt-5 gap-2 justify-self-center">
              {#each list_min_bet as rec}
                <div
                  on:click={() => {
                    handle_minbet(rec);
                  }} 
                  class="btn btn-sm btn-outline btn-success cursor-pointer">{new Intl.NumberFormat().format(rec)}</div>
              {/each}
            </div>
            
        </div>
    </div>
</div>

<input type="checkbox" id="my-modal-allinvoice" class="modal-toggle" bind:checked={isModal_allinvoice}>
<div class="modal" on:click|self={()=>isModal_allinvoice = false}>
  <div class="modal-box relative select-none  lg:max-w-7xl h-full lg:max-h-[600px] rounded-none lg:rounded-lg p-2 lg:p-4 overflow-hidden">
    <label for="my-modal-allinvoice" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
    <h3 class="text-xs lg:text-sm font-bold -mt-1">INVOICE</h3>
    <div class="overflow-auto h-[90%] scrollbar-thin scrollbar-thumb-green-100 mt-4">
        <table class="table table-xs w-full" >
            <thead class="sticky top-0">
                <tr>
                    <th width="1%" class="text-xs text-center align-top">ID</th>
                    <th width="4%" class="text-xs text-center align-top">DATE</th>
                    <th width="4%" class="text-xs text-center align-top">STATUS</th>
                    <th width="1%" class="text-xs text-right align-top">ROUND<br />BET</th>
                    <th width="5%" class="text-xs text-right align-top">BET</th>
                    <th width="5%" class="text-xs text-right align-top">CREDIT<br />BEFORE</th>
                    <th width="5%" class="text-xs text-right align-top">TOTAL<br />BET</th>
                    <th width="5%" class="text-xs text-right align-top">WIN</th>
                    <th width="5%" class="text-xs text-right align-top">CREDIT<br />AFTER</th>
                    <th width="*" class="text-xs text-left align-top">RESULT</th>
                </tr>
            </thead>
            <tbody>
                {#each list_datasend as rec}
                <tr>
                  <td class="text-xs  text-center whitespace-nowrap align-top">{rec.id_transaksi}</td>
                  <td class="text-xs  text-center whitespace-nowrap align-top">{rec.date_transaksi}</td>
                  <td class="text-xs  text-center whitespace-nowrap align-top">
                    <span class="{rec.status_transaction_css} p-1.5 text-xs lg:text-sm  uppercase  rounded-lg w-20 ">{rec.status_transaction}</span>
                  </td>
                  <td class="text-xs text-right link-accent whitespace-nowrap align-top">{new Intl.NumberFormat().format(rec.round_bet)}</td>
                  <td class="text-xs text-right text-error whitespace-nowrap align-top">(-{new Intl.NumberFormat().format(rec.bet)})</td>
                  <td class="text-xs text-right link-accent whitespace-nowrap align-top">{new Intl.NumberFormat().format(rec.credit_before)}</td>
                  <td class="text-xs text-right text-error whitespace-nowrap align-top">-{new Intl.NumberFormat().format(rec.total_bet)}</td>
                  <td class="text-xs text-right text-secondary whitespace-nowrap align-top">(+{new Intl.NumberFormat().format(rec.win)})</td>
                  <td class="text-xs text-right link-accent whitespace-nowrap align-top">{new Intl.NumberFormat().format(rec.credit_after)}</td>
                  <td class="text-xs text-left whitespace-nowrap w-52">
                    <div class="text-xs">
                      {rec.note_win}
                    </div>
                    <div class="grid grid-cols-7 w-full mt-5">
                      {#each rec.result_card as rec2}
                        <img  src="{rec2.img}" alt="" srcset=""> 
                      {/each}
                    </div>
                    <div class="grid grid-cols-7 mt-5 text-xs">
                      {#each rec.result_card_win as rec2}
                        <img  src="{rec2.img}" alt="" srcset=""> 
                      {/each}
                    </div>
                  </td>
                </tr>
                {/each}
            </tbody>
        </table>
    </div>
  </div>
</div>

<input type="checkbox" id="my-modal-carabermain" class="modal-toggle" bind:checked={isModal_carabermain}>
<div class="modal" on:click|self={()=>isModal_carabermain = false}>
  <div class="modal-box relative select-none  lg:max-w-4xl h-full lg:max-h-[600px] rounded-none lg:rounded-lg p-2 lg:p-4 overflow-hidden">
    <label for="my-modal-carabermain" class="btn btn-sm btn-circle absolute right-2 top-2">✕</label>
    <h3 class="text-xs lg:text-sm font-bold -mt-1">CARA BERMAIN</h3>
    <div class="overflow-auto h-[90%] scrollbar-thin scrollbar-thumb-green-100 mt-4">
        
    </div>
  </div>
</div>


<div class="hidden">
{#each card_result_data as rec}
  <img src="{rec.img}" alt="" srcset="">
{/each}
</div>
<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  
  .container {
    width: 100%;
  }
  @media (min-width: 640px) {
      .container {
          max-width: 640px;
      }
  }
  @media (min-width: 768px) {
      .container {
          max-width: 768px;
      }
  }
  @media (min-width: 1024px) {
      .container {
          max-width: 1024px;
      }
  }
  @media (min-width: 1280px) {
      .container {
          max-width: 1280px;
      }
  }
  @media (min-width: 1536px) {
      .container {
          max-width: 768px;
      }
  }
</style>
