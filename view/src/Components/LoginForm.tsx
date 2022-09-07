import {
  FormControl,
  FormLabel,
  FormErrorMessage,
  FormHelperText,
  Input,
  Button,
} from "@chakra-ui/react";
import { EventHandler, useState } from "react";

// TODO: validate form values
const LoginForm: React.FunctionComponent<any> = (props) => {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  let isError = email === "";

  const handleEmailChange = (e: any) => {
    setEmail(e.target.value);
  };
  const handlePasswordChange = (e: any) => setPassword(e.target.value);

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    console.log("hello");
    let reqbody: any = { user_id: email, password: password };

    // fetch("http://localhost:8080/v1/login", {
    //   method: "POST",
    //   body: reqbody,
    // })
    //   .then((res) => res.json())
    //   .then((res) => console.log(res))
    //   .catch((err) => console.error(err));
  };

  return (
    <form>
      <FormControl isInvalid={isError}>
        <FormLabel>Email</FormLabel>
        <Input
          type="email"
          value={email}
          onChange={handleEmailChange}
          isRequired
        />
        {!isError ? (
          <FormHelperText>
            Enter the email you'd like to receive the newsletter on.
          </FormHelperText>
        ) : (
          <FormErrorMessage>Email is required.</FormErrorMessage>
        )}
      </FormControl>
      <FormControl>
        <FormLabel>Password</FormLabel>
        <Input
          type="password"
          value={password}
          onChange={handlePasswordChange}
          isRequired
        />
      </FormControl>

      <Button
        type="submit"
        variant="outline"
        width="full"
        mt={4}
        onSubmit={(e) => handleSubmit(e)}
      >
        Sign In
      </Button>
    </form>
  );
};

export default LoginForm;
