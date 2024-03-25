import "./App.scss";

import Form from "./components/Form";
import Header from "./components/Header";

function App() {
  return (
    <>
      <Header />
      <main className="main">
        <div className="container">
          <Form />
        </div>
      </main>
    </>
  );
}

export default App;
