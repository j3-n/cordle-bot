const { initializeApp, applicationDefault, cert } = require('firebase-admin/app');
const { getFirestore, Timestamp, FieldValue } = require('firebase-admin/firestore');

class FirebaseConnect {
    constructor() {
        this.serviceAccount = undefined;
        this.database = undefined;
    }
    
    initialize() {
        this.serviceAccount = require('../fb-config/cordle-test-firebase-adminsdk-16gaq-61344483e9.json');

        initializeApp({
            credential: cert(this.serviceAccount)
        });

        this.database = getFirestore();
    }

    async getSnapshop() {
        const snapshot = await this.database.collection('users').get();
        snapshot.forEach((doc) => {
            console.log(doc.data().name, '=>', doc.data().id);
        });
    }

    async getCollection(collection) {
        const data = await this.database.collection(collection).get();
        const objects = [];

        data.forEach((obj) => {
            objects.push(obj);
        });

        return objects;
    }

    async getDocument(collection, document) {
        return await this.database.collection(collection).doc(document).get();
    }

    async updateCollection(undefined) {

    }

    async updateDocument(collection, document, object) {
        const docRef = this.database.collection(collection).doc(document);

        await docRef.update(object);
    }

    async addDocument(collection, document, object) {
        const docRef = this.database.collection(collection).doc(document);

        await docRef.set(object);
    }
}

module.exports.FirebaseConnect = FirebaseConnect;