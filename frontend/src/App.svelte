<script>
  import Home from "./pages/Home.svelte"

  export let path_api = "";
  export let version = "";
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);
  const token_browser = urlParams.get("token");
  
  if (token_browser === null) {
    console.log("TOKEN NOT FOUND");
  } else {
    initTimezone();
  }
  let flag_game = false;
  let client_credit = 0;
  let client_ipaddress = "0.0.0.0";
  let client_timezone = "";
  let client_company = "";
  let client_username = "";
  let client_name = "";
  let client_listbet = [];
  let loader_progress = 0
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

  async function initTimezone() {
    const res = await fetch(path_api+"api/healthz");
    if (!res.ok) {
      const message = `An error has occured: ${res.status}`;
      throw new Error(message);
    } else {
      const json = await res.json();
      client_ipaddress = json.real_ip;
      client_timezone = "Asia/Jakarta";
    }
    initapp(token_browser);
  }
  
  async function initapp(token) {
      const res = await fetch(path_api+"api/checktoken", {
          method: "POST",
          headers: {
              "Content-Type": "application/json",
          },
          body: JSON.stringify({
            client_token: token,
          }),
      });
      const json = await res.json();
      if (json.status === 400) {
          // logout();
      } else if (json.status == 403) {
          // alert(json.message);
      } else {
          flag_game = true;
          client_company = json.client_company;
          client_name = json.client_name;
          client_username = json.client_username;
          client_credit = json.client_credit;
          client_listbet = json.client_listbet.record;
      }
  }

  for(let i=0;i<card_result_data.length;i++){
    loader_progress = 0
    fetch(card_result_data[i].img)
    .then(response => {
      const contentLength = response.headers.get('content-length');
      let loaded = 0;
      return new Response(
        new ReadableStream({
          start(controller) {
            const reader = response.body.getReader();
            read();
            function read() {
              reader.read()
              .then((progressEvent) => {
                if (progressEvent.done) {
                  controller.close();
                  return; 
                }
                
                loaded += progressEvent.value.byteLength;
                //asdasd
                loader_progress = Math.round(loaded/contentLength*100);
                controller.enqueue(progressEvent.value);
                read();

              })
            }
          }
        })
      );
    })
    .then(response => response.blob()) // Read new readable stream to blob
    .then(blob => {
      const img = new Image()
      img.src = URL.createObjectURL(blob);
      document.body.appendChild(img);
      // Create new URL to blob image, set as src of image and append to DOM
    })
  }
</script>

  {#if flag_game &&  (parseInt(loader_progress) == 100)}
    <main class="container mx-auto px-2 text-base-content glass xl:rounded-box xl:mt-7 max-w-screen-xl bg-opacity-60 pb-5 h-fit lg:h-full">
      <Home
        {path_api} 
        {token_browser} 
        {client_company} 
        {client_timezone} 
        {client_ipaddress} 
        {client_username} 
        {client_name} 
        {client_credit} 
        {client_listbet} />
    </main>
  {:else}
  <main class="container mx-auto px-2 text-base-content glass xl:rounded-box xl:mt-7 max-w-screen-xl bg-opacity-60 pb-5 h-fit lg:h-[650px]"></main>
  {/if}

<footer class="footer footer-center p-4 text-base-content mt-2 text-center select-none">
  <div class="grid">
    <p class="text-xs lg:text-sm text-center">
      {version}
      <br />
      PowerBy
    </p>
    <img src="https://sdsb4d.com/logo-green.svg" alt="SDSB" class="w-24 lg:w-28">
  </div>
</footer>




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
