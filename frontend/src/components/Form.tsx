import "../styles/Form.scss";

const Form = () => {
  const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;
    if (!form) return;
    const formData = new FormData(form);
    const urlSearchParams = new URLSearchParams();

    formData.forEach((value, key) => {
      urlSearchParams.append(key, value.toString());
    });
    const response = await fetch("http://localhost:8080/data/", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: urlSearchParams,
    });
    if (response.ok) {
      const json = await response.json();
      console.log(json.status);
      form.reset();
    } else {
      console.error("Error");
    }
  };
  return (
    <div className="form-holder">
      <form className="form" onSubmit={handleSubmit}>
        <label htmlFor="name">Name</label>
        <input type="text" id="name" name="name" />
        <label htmlFor="email">Email</label>
        <input type="email" id="email" name="email" />
        <label htmlFor="message">Message</label>
        <textarea id="message" name="message" />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default Form;
