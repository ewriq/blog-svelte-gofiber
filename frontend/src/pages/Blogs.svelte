<script>
// @ts-nocheck

    export let title;
    import axios from "axios";
    let data = "";

    async function GetData() {
    if (title) {
      try {
        const response = await axios.post(
          `http://localhost:3000/api/v1/blog/view`, {
            title: title
          }
        );

        if (response.data.status === "OK") {
          data = response.data.data;
          console.log(response.data);
        }
        if (response.data.status === 404) {
            setTimeout(function () {
            window.location.href = "/404";
        }, 1);
        }
      } catch (err) {
        console.error(err);
      }
    } else {

    }
  }
  GetData();
</script>

<main>
    <div class="mt-5 ml-10">
        <div> 
            <h1 class="text-5xl">Başlık: {title}</h1>
        </div>
        <div>
            <h4>Açıklama kısmı: </h4>
            {@html data.description}
        </div>
    </div>
</main>