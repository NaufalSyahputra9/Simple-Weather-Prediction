document.getElementById('searchBtn').addEventListener('click', async () => {
    const city = document.getElementById('cityInput').value;
    const resultDiv = document.getElementById('result');

    if (!city) return alert("Isi dulu nama kotanya, Bud!");

    try {
        // Manggil backend Go yang jalan di port 8080
        const response = await fetch(`http://localhost:8080/weather/${city}`);
        const data = await response.json();

        if (response.ok) {
            document.getElementById('cityName').innerText = data.name;
            document.getElementById('temp').innerText = data.main.temp;
            document.getElementById('desc').innerText = data.weather[0].description;
            document.getElementById('humi').innerText = data.main.humidity;
            resultDiv.style.display = 'block';
        } else {
            alert("Kota tidak ditemukan!");
        }
    } catch (err) {
        console.error("Error nembak API:", err);
        alert("Server Backend-mu mati mungkin?");
    }
});