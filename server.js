import express from 'express'
import pool from './db.js'
import cors from 'cors'
const app = express()
const PORT = 3000

app.use(cors())

app.get('/citiry', async (req,res) =>{
    const lang = req.query.lang
    try{
       if(lang === 'ru'){
        const result = pool.query('SELECT name, id FROM citiry')
        res.json((await result).rows)
       } else{
        const result = pool.query('SELECT name, id FROM cities')
        res.json((await result).rows)
       }
    }
    catch (error){
        console.error('Ошибка при выполнении запроса', error);
        res.status(500).send('Ошибка сервера');
    }
    
})

app.get('/attracciti', async (req, res) => {
    const id = req.query.id
    console.log(id);
    
    try{
        const result = pool.query('SELECT * FROM attractions WHERE iden = $1', [id])
        res.json(result.rows)
    }catch(error){
        console.error('Ошибка при выполнении запроса', error);
        res.status(500).send('Ошибка сервера');
    }
   
})




app.listen(PORT, () => {
    console.log(`Сервер запущен на http://localhost:${PORT}`);
});