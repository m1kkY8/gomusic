const express = require("express");
const playdl = require("play-dl")

const app = express();
const port = 3001;

app.get("/search", async (req, res) => {
    const query = req.query.q

    if(!query){
        res.status(400).json({error: "No query"});
    }

    try {
        const results = await playdl.search(query, {limit: 1});
        
        const videoLinks = results.map(video => ({
            title: video.title,
            url: video.url
        }));
        res.json(videoLinks);
    }

    catch (error){
        console.log(error);
    }
});

app.listen(port);
