'use strict';

const {
	Contract
} = require('fabric-contract-api');

class GoldContract extends Contract {

	//Init function
	async InitGold() {
		console.log('Gold contract has been started');
	}

	//Add New Metal
	async AddMetal(ctx, args) {

		const data = JSON.parse(args);
		const metalId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'Metal');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].name == data.name) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. Metal Period with = ${data.name} already exists.`);
			}
		}
		let metal = {
			docType: 'Metal',
			id: metalId,
			name: data.name,
			icon: data.icon,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		console.log(JSON.stringify(metal));
		await ctx.stub.putState(metalId, Buffer.from(JSON.stringify(metal)));
		return metal;
	}

	//Add new Metal Group
	async AddMetalGroup(ctx, args) {

		const data = JSON.parse(args);
		const metalgroupId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'MetalGroup');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].name == data.name) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. MetalGroup with = ${data.name} already exists.`);
			}
		}
		//get metal data
		var metalAsBytes = await ctx.stub.getState(data.metalId);
		if (!metalAsBytes || metalAsBytes.length === 0) {
			throw new Error(`${data.metalId} does not exist`);
		}

		const metals = [];
		const result = JSON.parse(metalAsBytes.toString());
		metals.push(result);
		console.log(metals)
		let metalgroup = {
			docType: 'MetalGroup',
			id: metalgroupId,
			metals: metals,
			karatage: data.karatage,
			fineness: data.fineness,
			referenceId: data.referenceId,
			shortName: data.shortName,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		console.log(JSON.stringify(metalgroup));
		await ctx.stub.putState(metalgroupId, Buffer.from(JSON.stringify(metalgroup)));
		return metalgroup;
	}

	//Add new Cycle Period
	async AddCyclePeriod(ctx, args) {

		const data = JSON.parse(args);
		const cycleId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'CyclePeriod');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].name == data.name) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. Cycle Period with = ${data.name} already exists.`);
			}
		}
		let cycleperiod = {
			docType: 'CyclePeriod',
			id: cycleId,
			name: data.name,
			graceperiod: data.graceperiod,
			minWeight: data.minWeight,
			minValue: data.minValue,
			status: "Active",
			cycle: data.cycle,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		await ctx.stub.putState(cycleId, Buffer.from(JSON.stringify(cycleperiod)));
		return cycleperiod;
	}

	//Add new Product
	async AddProduct(ctx, args) {

		const data = JSON.parse(args);
		const productId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'Product');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].name == data.name) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. Cycle Period with = ${data.name} already exists.`);
			}
		}
		let product = {
			docType: 'Product',
			id: productId,
			name: data.name,
			images: data.images,
			video: data.video,
			status: "Active",
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		await ctx.stub.putState(productId, Buffer.from(JSON.stringify(product)));
		return product;
	}

	//Add new Collection
	async AddCollection(ctx, args) {

		const data = JSON.parse(args);
		const collectionId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'Collection');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].collection_name == data.collection_name) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. Cycle Period with = ${data.collection_name} already exists.`);
			}
		}
		let collection = {
			docType: 'Collection',
			id: collectionId,
			collection_name: data.collection_name,
			img1: data.img1,
			img2: data.img2,
			img3: data.img3,
			video: data.video,
			status: "Active",
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		await ctx.stub.putState(collectionId, Buffer.from(JSON.stringify(collection)));
		return collection;
	}

	//Add new Category
	async AddCategory(ctx, args) {

		const data = JSON.parse(args);
		const categoryId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'Category');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].category_name == data.category_name) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. Cycle Period with = ${data.category_name} already exists.`);
			}
		}
		let category = {
			docType: 'Category',
			id: categoryId,
			category_name: data.category_name,
			img1: data.img1,
			img2: data.img2,
			img3: data.img3,
			video: data.video,
			status: "Active",
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		await ctx.stub.putState(categoryId, Buffer.from(JSON.stringify(category)));
		return category;
	}

	//Add new Variety
	async AddVariety(ctx, args) {

		const data = JSON.parse(args);
		const varietyId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'Variety');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].variety_name == data.variety_name) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. Cycle Period with = ${data.variety_name} already exists.`);
			}
		}
		let variety = {
			docType: 'Variety',
			id: varietyId,
			variety_name: data.variety_name,
			img1: data.img1,
			img2: data.img2,
			img3: data.img3,
			video: data.video,
			status: "Active",
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		await ctx.stub.putState(varietyId, Buffer.from(JSON.stringify(variety)));
		return variety;
	}

	//Add new Charges
	async AddCharges(ctx, args) {

		const data = JSON.parse(args);
		const chargeId = data.id;
		console.log("incoming asset fields: ", data);

		// Retrieve the current charge using key provided
		var recordAsBytes = await ctx.stub.getState(chargeId);

		if (!recordAsBytes || recordAsBytes.length === 0) {
			throw new Error(`Error Message from charge: Order with orderId = ${chargeId} does not exist.`);
		}

		//get metal data
		var userAsBytes = await ctx.stub.getState(data.deleveryAgent);
		if (!userAsBytes || userAsBytes.length === 0) {
			throw new Error(`${data.deleveryAgent} does not exist`);
		}

		const deleveryAgent = [];
		const result = JSON.parse(userAsBytes.toString());
		deleveryAgent.push(result);

		let charges = {
			docType: 'Charges',
			id: chargeId,
			paymentId: data.paymentId,
			mode: data.mode,
			amount: data.amount,
			weight: data.weight,
			deleveryAgent: deleveryAgent,
			status: data.status,
			instantGoldAppiled: data.instantGoldAppiled,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		await ctx.stub.putState(chargeId, Buffer.from(JSON.stringify(charges)));
		return charges;
	}

	//Add New Calculation
	async AddCalculation(ctx, args) {

		const data = JSON.parse(args);
		const calId = data.id;
		console.log("incoming asset fields: ", data);

		let calculation = {
			docType: 'Calculation',
			id: calId,
			Type: data.Type,
			Percentage: data.Percentage,
			Status: "Active",
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		console.log(JSON.stringify(calculation));
		await ctx.stub.putState(calId, Buffer.from(JSON.stringify(calculation)));
		return calculation;
	}

	//Add new Plan
	async AddPlan(ctx, args) {

		const data = JSON.parse(args);
		const planId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'Plan');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].name == data.name && records[i].planType == data.planType) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. Plan with = ${data.name} and ${data.planType} already exists.`);
			}
		}
		//get cycleperiod data
		var cycleAsBytes = await ctx.stub.getState(data.cyclePeriod);
		if (!cycleAsBytes || cycleAsBytes.length === 0) {
			throw new Error(`${data.cyclePeriod} does not exist`);
		}

		//const cyclePeriod = [];
		const cyclePeriod = JSON.parse(cycleAsBytes.toString());
		//cyclePeriod.push(result);

		let plan = {
			docType: 'Plan',
			id: planId,
			name: data.name,
			cyclePeriod: cyclePeriod,
			mode: data.mode,
			bonus: data.bonus,
			planType: data.planType,
			duration: data.duration,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		console.log(JSON.stringify(plan));
		await ctx.stub.putState(planId, Buffer.from(JSON.stringify(plan)));
		return plan;
	}

	//Add Buy/Sell
	async AddBuySell(ctx, args) {
		const data = JSON.parse(args);
		const buysellId = data.id;
		console.log("incoming asset fields: ", data);

		let buysell = {
			docType: 'BuySell',
			id: buysellId,
			buy: data.buy,
			sell: data.sell,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		console.log(JSON.stringify(buysell));
		await ctx.stub.putState(buysellId, Buffer.from(JSON.stringify(buysell)));
		return buysell;
	}

	//Add Video
	async AddVideo(ctx, args){
		const data = JSON.parse(args);
		const videoId = data.id;
		console.log("incoming asset fields: ", data);

		let video = {
			docType: 'Video',
			id: videoId,
			language: data.language,
			category: data.category,
			video: data.video,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		console.log(JSON.stringify(video));
		await ctx.stub.putState(videoId, Buffer.from(JSON.stringify(video)));
		return video;
	}

	//Add User
	async AddUser(ctx, args){
		const data = JSON.parse(args);
		const userId = data.id;
		console.log("incoming asset fields: ", data);

		// Check if record already exits
		var recordAsBytes = await this.GetAll(ctx, 'User');
		var records = JSON.parse(recordAsBytes.toString());
		for (var i = 0; i < records.length; i++) {
			if (records[i].mobile == data.mobile) {
				console.log(records.length, i, "for loop")
				throw new Error(`Error Message. Metal Period with = ${data.mobile} already exists.`);
			}
		}

		let user = {
			docType: 'User',
			id: userId,
			fname: data.fname,
			email: data.email,
			mobile: data.mobile,
			dob: data.dob,
			pan: data.pan,
			isWhatsapp: data.isWhatsapp,
			role: data.role,
			addresses: [],
			isInvested: data.isInvested || false,
			image: data.image,
			referral: data.referral,
			referenceType: data.referenceType,
			refCode: data.refCode || '',
			GBPcode: data.GBPcode || '',
			referralBonusEntries: data.referralBonusEntries || [],
			GBPBonusEntries: data.GBPBonusEntries || [],
			joiningBonus: data.joiningBonus || 0,
			level: data.level,
			deviceToken: data.deviceToken || '',
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		};

		// === Save asset to state ===
		console.log(JSON.stringify(user));
		await ctx.stub.putState(userId, Buffer.from(JSON.stringify(user)));
		return user;
	}

	//Add Installment
	async AddInstallment(ctx, args){
		const data = JSON.parse(args);
		const installmentId = data.id;
		console.log("incoming asset fields: ", data);

		let installment = {
			docType: "Installment",
			id: data.id,
			paymentId: data.paymentId,
			user: data.user,
			mode: data.user || "online",
			status: data.status,
			amount: data.amount || 0,
			gold: data.gold || 0,
			collector: data.collector,
			otp: data.otp || '',
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

			// === Save asset to state ===
			await ctx.stub.putState(installmentId, Buffer.from(JSON.stringify(installment)));
			return installment;
	}
	//Add Subscription
	async AddSubscription(ctx, args){
		const data = JSON.parse(args);
		const subscriptionId = data.id;
		console.log("incoming asset fields: ", data);

		let subscription = {
			docType: "Subscription",
			id: subscriptionId,
			user: data.user,
			status: data.status || "Processing",
			address: data.address,
			plan: data.plan,
			installments: data.installments || [],
			maturityDate: data.maturityDate || null,
			planBonus: data.planBonus,
			unpaidSkips: data.unpaidSkips || 0,
			skipCount: data.skipCount || 0,
			unpaidInvestments: data.unpaidInvestments || 0,
			trackingId: data.trackingId || '',
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()

		}


		// === Save asset to state ===
		console.log(JSON.stringify(subscription));
		await ctx.stub.putState(subscriptionId, Buffer.from(JSON.stringify(subscription)));
		return subscription;

	}

	//Add Referral Bonus
	async AddReferralBonus(ctx, args){
		const data = JSON.parse(args);
		const referralId = data.id;

		let referral = {
			docType: "ReferralBonus",
			id: referralId,
			user: data.user,
			refereedBy: data.refereedBy,
			subscription: data.subscription,
			bonus: data.bonus,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(referral));
		await ctx.stub.putState(referralId, Buffer.from(JSON.stringify(referral)));
		return referral;
	}

	//Add Address
	async AddAddress(ctx, args){
		const data = JSON.parse(args);
		const addressId = data.id;

		//get user data
		var userAsBytes = await ctx.stub.getState(data.user);
		if (!userAsBytes || userAsBytes.length === 0) {
			throw new Error(`${data.user} does not exist`);
		}

		let userData = JSON.parse(userAsBytes.toString());
		const userdetails = {
			fname: userData.fname,
			email: userData.email,
			id: userData.id,
			mobile: userData.mobile
		}
		console.log(userData, "User Data")

		let address = {
			docType: "Address",
			id: addressId,
			user: userdetails,
			addressType: data.addressType || "Home",
			isDefaultAddress: data.isDefaultAddress || false,
			status: data.status,
			pin: data.pin,
			landMark: data.landMark,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(address));
		await ctx.stub.putState(addressId, Buffer.from(JSON.stringify(address)));
		let addresses = [];
		userData.addresses = [...addresses, address]
		console.log(addresses, userData)
		await ctx.stub.putState(data.user, Buffer.from(JSON.stringify(userData)));

		return address;
	}

	//Add Bank
	async AddBank(ctx, args){
		const data = JSON.parse(args);
		const bankId = data.id;

		//get user data
		var userAsBytes = await ctx.stub.getState(data.user);
		if (!userAsBytes || userAsBytes.length === 0) {
			throw new Error(`${data.user} does not exist`);
		}

		let userData = JSON.parse(userAsBytes.toString());
		console.log(userData, "User Data")

		let bank = {
			docType: "Bank",
			id: bankId,
			user: data.user,
			Accountnum: data.Accountnum,
			IFSC: data.IFSC,
			Bank: data.Bank,
			Branch: data.Branch,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(bank));
		await ctx.stub.putState(bankId, Buffer.from(JSON.stringify(bank)));
		userData.banks = [...banks, bank]
		console.log(addresses, userData)
		await ctx.stub.putState(data.user, Buffer.from(JSON.stringify(userData)));

		return bank;
	}

	//Create Oppointment
	async CreateOppointment(ctx, args){
		const data = JSON.parse(args);
		const oppointmentId = data.id;

		//get user data
		var userAsBytes = await ctx.stub.getState(data.user);
		if (!userAsBytes || userAsBytes.length === 0) {
			throw new Error(`${data.user} does not exist`);
		}

		let userData = JSON.parse(userAsBytes.toString());
		const userdetails = {
			fname: userData.fname,
			email: userData.email,
			id: userData.id,
			mobile: userData.mobile
		}
		console.log(userData, "User Data")

		//get metalgroup data
		var mgAsBytes = await ctx.stub.getState(data.metalGroup);
		if (!mgAsBytes || mgAsBytes.length === 0) {
			throw new Error(`${data.metalGroup} does not exist`);
		}

		let mgData = JSON.parse(mgAsBytes.toString());

		//get buysell price
		var buysellAsBytes = await ctx.stub.getState(data.buySellPrice)
		if (!buysellAsBytes || buysellAsBytes.length === 0) {
			throw new Error(`${data.buySellPrice} does not exist`);
		}

		let buysellData = JSON.parse(buysellAsBytes.toString());

		//get verifier data
		var verifierAsBytes = await ctx.stub.getState(data.verifier);
		if (!verifierAsBytes || verifierAsBytes.length === 0) {
			throw new Error(`${data.verifier} does not exist`);
		}

		let verifierData = JSON.parse(verifierAsBytes.toString());

		let oppointment = {
			docType: "Oppointment",
			id: oppointmentId,
			user: userdetails,
			weight: data.weight,
			metalGroup: mgData,
			buySellPrice: buysellData,
			otp: data.otp,
			verifier: verifierData,
			status: data.status,
			appointmentTime: data.appointmentTime,
			appointmentDate: data.appointmentDate,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(oppointment));
		await ctx.stub.putState(oppointmentId, Buffer.from(JSON.stringify(oppointment)));

		return oppointment;
	}

	//Add Item
	async AddItem(ctx, args){
		const data = JSON.parse(args);
		const itemId = data.id;

		let item = {
			docType: "Item",
			id: itemId,
			name: data.name,
			images: data.images,
			video: data.video,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(item));
		await ctx.stub.putState(itemId, Buffer.from(JSON.stringify(item)));
		return item;
	}

	//Add Item Details
	async AddItemdetails(ctx, args){

	}

	//Add Level
	async AddLevel(ctx, args){
		const data = JSON.parse(args);
		const levelId = data.id;

		let level = {
			docType: "Level",
			id: levelId,
			name: data.name,
			commission: data.commission,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(level));
		await ctx.stub.putState(levelId, Buffer.from(JSON.stringify(level)));
		return level;
	}

	//Add permission
	async AddPermission(ctx, args){
		const data = JSON.parse(args);
		const permissionId = data.id;

		let permission = {
			docType: "Permission",
			id: permissionId,
			permission_name: data.permission_name,
			status: data.status || "active",
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(permission));
		await ctx.stub.putState(permissionId, Buffer.from(JSON.stringify(permission)));
		return permission;
	}

	//Add Pincode
	async AddPincode(ctx, args){
		const data = JSON.parse(args);
		const pincodeId = data.id;

		let pincode = {
			docType: "Pincode",
			id: pincodeId,
			pin_number: data.pin_number,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(pincode));
		await ctx.stub.putState(pincodeId, Buffer.from(JSON.stringify(pincode)));
		return pincode;
	}
	//Create Order
	async CreateOrder(ctx, args){
		const data = JSON.parse(args);
		const orderId = data.id;

		let order = {
			docType: "Order",
			id: orderId,
			user: data.user,
			cart: data.cart,
			transactions: data.transactions,
			status: data.status || "Processing",
			address: data.address,
			consignment: data.consignment || '',
			deliveryCharge: data.deliveryCharge,
			buySell: data.buySell,
			redeemGoldApplied: data.redeemGoldApplied || 0,
			instantGoldApplied: data.instantGoldApplied || 0,
			otp: data.otp || '',
			subscriptionApplied: data.subscriptionApplied,
			totalCharges: data.totalCharges || 0,
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(order));
		await ctx.stub.putState(orderId, Buffer.from(JSON.stringify(order)));
		return order;
	}

	//Create Order Type
	async CreateOrderType(ctx, args){
		const data = JSON.parse(args);
		const orderTypeId = data.id;

		let ordertype = {
			docType: "Ordertype",
			id: orderTypeId,
			name: data.name,
			status: data.status || "Inactive",
			createdAt: ctx.stub.getDateTimestamp(),
			updatedAt: ctx.stub.getDateTimestamp()
		}

		//=== Save asset to state ===
		console.log(JSON.stringify(ordertype));
		await ctx.stub.putState(orderTypeId, Buffer.from(JSON.stringify(ordertype)));
		return ordertype;
	}

	//Update user
	async UpdateUser(ctx, args){
		const data = JSON.parse(args);
		
		let userAsBytes = await ctx.stub.getState(data.id)
		if (!userAsBytes || userAsBytes.length === 0) {
            throw new Error(`${data.id} does not exist`);
        }
        const user = JSON.parse(userAsBytes.toString());
		user.isInvested = data.isInvested;
		user.joiningBonus = data.joiningBonus;

        await ctx.stub.putState(data.id, Buffer.from(JSON.stringify(user)));
        console.info('============= END : Update User ===========');
	}
//---------------------------------------------------------------------------------------------
	//Get by Id
	async GetbyId(ctx, id, docType) {
		let queryString = {};
		queryString.selector = {};
		queryString.selector.docType = docType;
		queryString.selector.id = id;
		return await this.GetQueryResultForQueryString(ctx, JSON.stringify(queryString)); //shim.success(queryResults);
	}

	//Get All Records
	async GetAll(ctx, docType) {

		let queryString = {};
		queryString.selector = {};
		queryString.selector.docType = docType;
		console.log(queryString, "line 127");
		return await this.GetQueryResultForQueryString(ctx, JSON.stringify(queryString));
	}

	async GetQueryResultForQueryString(ctx, queryString) {

		let resultsIterator = await ctx.stub.getQueryResult(queryString);
		let results = await this.GetAllResults(resultsIterator, false);

		return JSON.stringify(results);
	}

	async GetAllResults(iterator, isHistory) {
		let allResults = [];
		let res = await iterator.next();
		while (!res.done) {
			if (res.value && res.value.value.toString()) {
				let jsonRes = {};
				console.log(res.value.value.toString('utf8'));
				if (isHistory && isHistory === true) {
					jsonRes.TxId = res.value.tx_id;
					jsonRes.Timestamp = res.value.timestamp;
					try {
						jsonRes.Value = JSON.parse(res.value.value.toString('utf8'));
					} catch (err) {
						console.log(err);
						jsonRes.Value = res.value.value.toString('utf8');
					}
				} else {
					jsonRes = JSON.parse(res.value.value.toString('utf8'));
					// jsonRes.Key = res.value.key;
					// try {
					// 	jsonRes.Record = JSON.parse(res.value.value.toString('utf8'));
					// } catch (err) {
					// 	console.log(err);
					// 	jsonRes.Record = res.value.value.toString('utf8');
					// }
				}
				allResults.push(jsonRes);
			}
			res = await iterator.next();
		}
		iterator.close();
		return allResults;
	}

	//Delete record from world state
	async DeleteAsset(ctx, id){
		console.info('============= deleteRecord ===========');
        if (id.length < 1) {
            throw new Error('Asset Id required as input')
        }
        console.log("Id = " + id);

		await ctx.stub.deleteState(id); //remove the asset from chaincode state
	}
}


module.exports = GoldContract;