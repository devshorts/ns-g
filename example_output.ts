/* Do not change, this code is generated from Golang structs */


export class Bank {
    name: string;

    static createFrom(source: any) {
        if ('string' === typeof source) source = JSON.parse(source);
        const result = new Bank();
        result.name = source["name"];
        return result;
    }

    //[Bank:]


    //[end]
}