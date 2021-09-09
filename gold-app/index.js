const express = require('express');
const Joi = require('joi');

const enrollGoldAdmin = require('./app/enrollAdmin');
const registerGoldUser = require('./app/registerUser');


const app = express();
app.use(express.json());


app.post('/api/:org/enroll/admin', async (req,res)=>{

    if(req.params.org != 'gold'){
        return res.status(400).send("This request was meant for gold organization");
    }

    let response = await enrollGoldAdmin();

    if(response && response.success){

        console.log(`Enroll was Success: ${response.message}`);
        res.status(200).json(response);

    }else{
        console.log(`Enroll was Failure: ${response.message}`);
        res.status(401).json(response);
        return;
    }
});


app.post('/api/:org/:user', async (req,res)=>{

    const schema = Joi.object({
        userAffiliation: Joi.string().min(3).required()
    });
    
    const result = schema.validate(req.body);
    
    console.log(result);
    
    if (result.error) {
        return res.status(400).send(result.error);
    }

    if(req.params.org != 'gold'){
        return res.status(400).send("This request was meant for gold organization");
    }

    if(!req.params.user && req.params.user.length >3 ){
        return res.status(400).send("Input validation error w.r.t user");
    }

    let response = await registerGoldUser(req.params.user,req.body.userAffiliation,req.params.org);

    if(response && response.success){

        console.log(`Enroll was Success: ${response.message}`);
        res.status(200).json(response);

    }else{
        console.log(`Enroll was Failure: ${response.message}`);
        res.status(401).json(response);
        return;
    }

    // if(req.body.userAffiliation == 'hdfc.user'){
        
    //     response = await registerUserOnApp(req.params.user, req.body.userAddress);
        
    //     if(response && response.success){
            
    //         console.log(`User Registration on App was Success: ${response.message}`);
            
    //         res.status(200).json(response);
    //     }else{
    //         console.log(`User Registration on App was Failure: ${response.message}`);
    //         res.status(401).json(response);
    //     }
    // }
});

let port = process.env.PORT || 3000;
app.listen(port, () => console.log(`Gold server listening on port ${port}....`));