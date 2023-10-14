async fn fetch_data_from_db() -> String {
    // ...
}

async fn render_data() {
    let data = fetch_data_from_db().await;

    println!("{data}");
}

trait Database {
    async fn fetch_data(&self) -> String;
}

impl Database for MyDB {
    async fn fetch_data(&self) -> String {
        // ...
    }
}

trait Database {
    type FetchData<'a>: Future<Output = String> + 'a where Self: 'a;
    fn fetch_data(&self) -> FetchData<'a>;
}