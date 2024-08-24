import Framily from "./assets/framily.jpg";
import "./App.css";

function App() {
  return (
    <>
      <div className="image-container">
        <img src={Framily} alt="Framily" className="centered-image" />
      </div>
      <div>Welcome to Ian Myjer Dot Com!</div>
    </>
  );
}

export default App;
