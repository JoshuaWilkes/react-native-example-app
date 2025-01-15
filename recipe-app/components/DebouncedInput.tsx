import { useState, useEffect } from "react";
import { TextInput, StyleSheet } from "react-native";

interface DebouncedInputProps {
  value: string;
  onChange: (value: string) => void;
  debounce?: number;
}

const DebouncedInput: React.FC<DebouncedInputProps> = ({
  value: initialValue,
  onChange,
  debounce = 250,
  ...rest
}) => {
  const [value, setValue] = useState(initialValue);

  useEffect(() => {
    setValue(initialValue);
  }, [initialValue]);

  useEffect(() => {
    const timeout = setTimeout(() => {
      onChange(value);
    }, debounce);

    return () => clearTimeout(timeout);
  }, [value]);

  return (
    <TextInput style={styles.textInput} value={value} onChangeText={setValue} />
  );
};

const styles = StyleSheet.create({
  textInput: {
    borderWidth: 1,
    borderColor: "#ebebeb",
    padding: 10,
    marginVertical: 5,
    color: "white",
  },
});

export default DebouncedInput;
