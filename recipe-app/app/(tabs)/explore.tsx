import { StyleSheet, Image, Platform, TextInput } from "react-native";

import { Collapsible } from "@/components/Collapsible";
import { ExternalLink } from "@/components/ExternalLink";
import ParallaxScrollView from "@/components/ParallaxScrollView";
import { ThemedText } from "@/components/ThemedText";
import { ThemedView } from "@/components/ThemedView";
import { IconSymbol } from "@/components/ui/IconSymbol";
import { useInfiniteQuery } from "@connectrpc/connect-query";
import { queryRecipes } from "../../gen/recipe/v1alpha1/recipe-RecipeService_connectquery";
import { useMemo, useState } from "react";
import Recipe from "../../components/Recipe";
import DebouncedInput from "../../components/DebouncedInput";

export default function TabTwoScreen() {
  const [search, setSearch] = useState<string>("");

  const { data, error, dataUpdatedAt } = useInfiniteQuery(
    queryRecipes,
    { pageToken: "", search },
    {
      getNextPageParam: (p) =>
        p.nextPageToken === "" ? undefined : p.nextPageToken,
      pageParamKey: "pageToken",
    }
  );

  const recipes = useMemo(() => {
    return data?.pages.flatMap((p) => p.recipes);
  }, [dataUpdatedAt]);

  return (
    <ParallaxScrollView
      headerBackgroundColor={{ light: "#D0D0D0", dark: "#353636" }}
      headerImage={
        <IconSymbol
          size={310}
          color="#808080"
          name="chevron.left.forwardslash.chevron.right"
          style={styles.headerImage}
        />
      }
    >
      <ThemedView style={styles.titleContainer}>
        <ThemedText type="title">Explore Recipes</ThemedText>
      </ThemedView>

      <DebouncedInput debounce={500} value={search} onChange={setSearch} />

      {recipes?.map((recipe, i) => {
        return <Recipe key={`rec-${recipe.id}`} recipe={recipe} />;
      })}
    </ParallaxScrollView>
  );
}

const styles = StyleSheet.create({
  headerImage: {
    color: "#808080",
    bottom: -90,
    left: -35,
    position: "absolute",
  },
  titleContainer: {
    flexDirection: "row",
    gap: 8,
  },
  textInput: {
    borderWidth: 1,
    borderColor: "#ebebeb",
    padding: 10,
    marginVertical: 5,
  },
});
