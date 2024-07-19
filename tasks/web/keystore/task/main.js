const express = require('express')
const bodyParser = require('body-parser')

const cookieParser = require('cookie-parser')
const mysql = require('mysql-await')

const app = express()
app.set('view engine', 'ejs')
app.use(cookieParser('321618637361347ba36458f1e7cb9acb8ee33fb050354d4c5a767ecee04470cf222d1bb815f55a7f8c84ac8b32a00cdc126b75c0c689026f21b5c46c8ee9e24c3ab367fa9b412f9d0f166'))
app.use(bodyParser.urlencoded({ extended: false }))

const conn = mysql.createPool({
    connectionLimit: 10,
    host     : 'mysql',
    user     : 'root',
    password : 'keystore',
    database : 'keystore'
})

app.get('/', async (req, res) => {
    if(!req.signedCookies.login || !req.signedCookies.password) {
        return unathorized(req, res)
    }
    const db = await conn.awaitGetConnection()
    try {
        accounts = await db.awaitQuery("SELECT `id` FROM `users` WHERE `login` = ? AND `password` = ? LIMIT 1", [req.signedCookies.login, req.signedCookies.password])
        if (accounts.length === 1) {
            privkeys = await db.awaitQuery("SELECT `privkey` FROM `privkeys` WHERE `id` = ? LIMIT 1", [accounts[0].id])
            if (privkeys.length === 1) {
                res.render('index', { privkey: privkeys[0].privkey })
            }else{
                res.render('index', { privkey: "Key is empty" })
            }
        } else {
            return unathorized(req, res)
        }
    } catch (e) {
        console.log(e)
    } finally {
        db.release()
    }
})

app.post('/privkey', async (req, res) => {
    if (req.body.privkey.length > 1024) {
        return unathorized(req, res)
    }
    const db = await conn.awaitGetConnection()
    try{
        accounts = await db.awaitQuery("SELECT `id` FROM `users` WHERE `login` = ? AND `password` = ? LIMIT 1", [req.signedCookies.login, req.signedCookies.password])
        if (accounts.length === 1) {
            keys = await db.awaitQuery("SELECT `id` FROM `privkeys` WHERE `id` = ?", [accounts[0].id])
            if (keys.length === 0) {
                await db.awaitQuery("INSERT INTO `privkeys` (`id`, `privkey`) VALUES (?, ?)", [accounts[0].id, req.body.privkey])
            }else{
                await db.awaitQuery("UPDATE `privkeys` SET `privkey` = ? WHERE id = ?", [req.body.privkey, accounts[0].id])
            }
        }else{
            console.log("unauthorized")
            return unathorized(req, res)
        }
        return res.redirect('/')
    } catch(e) {
        console.log(e)
    } finally {
        db.release()
    }
})


app.get('/login', (req, res) => {
    res.render('login')
})

app.post('/login', async (req, res) => {
    if (req.body.login.length > 250 || req.body.password.length > 250) {
        return unathorized(req, res)
    }
    const db = await conn.awaitGetConnection()
    try {
        result = await db.awaitQuery("SELECT `id` FROM `users` WHERE `login` = ? AND `password` = ? LIMIT 1", [req.body.login, req.body.password])
        if (result.length === 1) {
            res.cookie('login', req.body.login, { signed: true })
            res.cookie('password', req.body.password, { signed: true })
            return res.redirect('/')
        } else {
            return unathorized(req, res)
        }
    } catch(e) {
        console.log(e)
    } finally {
        db.release()
    }
})

app.get('/signup', (req, res) => {
    res.render('signup')
})

app.post('/signup', async (req, res) => {
    if (req.body.login.length > 250 || req.body.password.length > 250) {
        return unathorized(req, res)
    }
    const db = await conn.awaitGetConnection()
    try {
        result = await db.awaitQuery("SELECT `id` FROM `users` WHERE `login` = ?", [req.body.login])
        if (result.length !== 0)
            return unathorized(req, res)
        result = await db.awaitQuery("INSERT INTO `users` (`login`, `password`) VALUES (?, ?)", [req.body.login, req.body.password])
        res.cookie('login', req.body.login, { signed: true })
        res.cookie('password', req.body.password, { signed: true })
        return res.redirect('/')
    } catch(e) {
        console.log(e)
    } finally {
        db.release()
    }
})

app.get('/logout', (req, res) => {
    return unathorized(req, res)
})

app.listen(5010, () => {
    console.log(`Task started`)
})

function unathorized(req, res) {
    res.clearCookie('login')
    res.clearCookie('password')
    return res.redirect('/login')
}
