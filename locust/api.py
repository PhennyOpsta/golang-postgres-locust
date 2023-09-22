
class Api:
    def __init__(self, client):
        self.client = client

    def get_text_root(self):
        self.client.get("/")

    def get_text_api(self):
        self.client.get("/api")

    def post_plus(self, first_value, second_value):
        result = self.client.post("/api/plus", json={
            "a": first_value,
            "b": second_value,
        })
        return result

    def post_minus(self, first_value, second_value):
        result = self.client.post("/api/minus", json={
            "a": first_value,
            "b": second_value,
        })
        return result

    def post_multiple(self, first_value, second_value):
        result = self.client.post("/api/multiple", json={
            "a": first_value,
            "b": second_value,
        })
        return result

    def post_divide(self, first_value, second_value):
        result = self.client.post("/api/divide", json={
            "a": first_value,
            "b": second_value,
        })
        return result

    def post_product(self, product_id):
        result = self.client.post("/api/product/" + product_id)
        return result
